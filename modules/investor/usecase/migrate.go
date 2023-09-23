package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/rzfhlv/doit/modules/investor/model"
	logrus "github.com/rzfhlv/doit/utilities/log"
)

var (
	INVESTOR = "INVESTOR"
	PENDING  = "PENDING"
	FAILED   = "FAILED"

	UPSERTMONGOLOG  = "Investor Usecase UpsertMongo"
	MARSHALLOG      = "Investor Usecase upsertInvestor Marshal"
	GETINVESTORSLOG = "Investor Usecase getInvestors"
	DATAMIGRATED    = "Data Migrated"
)

func (u *Usecase) ConventionalMigrate(ctx context.Context) (err error) {
	investors, err := u.repo.GetPsql(ctx)
	if err != nil {
		return
	}
	for _, investor := range investors {
		err = u.repo.UpsertMongo(ctx, investor)
		if err != nil {
			return
		}
	}
	return
}

func (u *Usecase) MigrateInvestors(ctx context.Context) (err error) {
	chInvestor := u.getInvestors()
	workerInvestor1 := u.upsertInvestors(chInvestor)
	workerInvestor2 := u.upsertInvestors(chInvestor)
	workerInvestor3 := u.upsertInvestors(chInvestor)
	workerInvestor4 := u.upsertInvestors(chInvestor)
	workerInvestor5 := u.upsertInvestors(chInvestor)
	chanInvestorSum := u.mergeChanInvestor(workerInvestor1, workerInvestor2, workerInvestor3, workerInvestor4, workerInvestor5)

	// print output
	counterTotal := 0
	for investor := range chanInvestorSum {
		if investor.ID > 0 {
			counterTotal++
		}
	}
	logrus.Log(nil).Info(fmt.Sprintf("%d "+DATAMIGRATED, counterTotal))

	return
}

func (u *Usecase) getInvestors() <-chan model.Investor {
	chanOut := make(chan model.Investor)

	go func() {
		investors, err := u.repo.GetPsql(context.Background())
		if err != nil {
			logrus.Log(nil).Error(fmt.Sprintf(GETINVESTORSLOG+" %v", err.Error()))
			return
		}
		for _, investor := range investors {
			chanOut <- investor
		}

		close(chanOut)
	}()

	return chanOut
}

func (u *Usecase) upsertInvestors(chanIn <-chan model.Investor) <-chan model.Investor {
	chanOut := make(chan model.Investor)

	go func() {
		ctx := context.Background()
		for investor := range chanIn {
			now := time.Now()
			payload, err := json.Marshal(investor)
			if err != nil {
				logrus.Log(nil).Error(fmt.Sprintf(MARSHALLOG+" %v", err.Error()))
				return
			}

			// save to outbox table
			err = u.saveOutbox(ctx, now, payload, investor)
			if err != nil {
				return
			}

			// migrations
			err = u.upsertMongo(ctx, now, payload, investor)
			if err != nil {
				return
			}

			// delete outbox
			err = u.deleteOutbox(ctx, investor.ID)
			if err != nil {
				return
			}
			chanOut <- investor
		}
		close(chanOut)
	}()

	return chanOut
}

func (u *Usecase) mergeChanInvestor(chanInMany ...<-chan model.Investor) <-chan model.Investor {
	wg := new(sync.WaitGroup)
	chanOut := make(chan model.Investor)

	wg.Add(len(chanInMany))
	for _, eachChan := range chanInMany {
		go func(eachChan <-chan model.Investor) {
			for eachChanData := range eachChan {
				chanOut <- eachChanData
			}
			wg.Done()
		}(eachChan)
	}

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}

func (u *Usecase) saveOutbox(ctx context.Context, now time.Time, payload []byte, investor model.Investor) (err error) {
	outBox := model.Outbox{
		Identifier: investor.ID,
		Payload:    string(payload),
		Event:      INVESTOR,
		Status:     PENDING,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	err = u.repo.UpsertOutbox(ctx, outBox)
	return
}

func (u *Usecase) upsertMongo(ctx context.Context, now time.Time, payload []byte, investor model.Investor) (err error) {
	err = u.repo.UpsertMongo(ctx, investor)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf(UPSERTMONGOLOG+" %v", err.Error()))
		outBox := model.Outbox{
			Identifier: investor.ID,
			Payload:    string(payload),
			Event:      INVESTOR,
			Status:     FAILED,
			CreatedAt:  now,
			UpdatedAt:  now,
		}
		err = u.repo.UpsertOutbox(ctx, outBox)
		return
	}
	return
}

func (u *Usecase) deleteOutbox(ctx context.Context, id int64) (err error) {
	err = u.repo.DeleteOutbox(ctx, id)
	return
}
