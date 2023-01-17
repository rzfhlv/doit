package usecase

import (
	"context"
	"doit/modules/investor/model"
	"doit/modules/investor/repository"
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"
)

type IUsecase interface {
	MigrateInvestors(ctx context.Context) error
	ConventionalMigrate(ctx context.Context) error
}

type Usecase struct {
	repo repository.IRepository
}

func NewUsecase(repo repository.IRepository) IUsecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) ConventionalMigrate(ctx context.Context) error {
	investors, err := u.repo.GetPsql(ctx)
	if err != nil {
		log.Printf("error get psql")
		return err
	}
	for _, investor := range investors {
		err = u.repo.UpsertMongo(ctx, investor)
		if err != nil {
			log.Printf("error update mongo")
			return err
		}
	}
	return nil
}

func (u *Usecase) MigrateInvestors(ctx context.Context) error {
	chInvestor := u.getInvestors()
	workerInvestor1 := u.upsertInvestors(chInvestor)
	workerInvestor2 := u.upsertInvestors(chInvestor)
	workerInvestor3 := u.upsertInvestors(chInvestor)
	workerInvestor4 := u.upsertInvestors(chInvestor)
	workerInvestor5 := u.upsertInvestors(chInvestor)
	chanInvestorSum := mergeChanInvestor(workerInvestor1, workerInvestor2, workerInvestor3, workerInvestor4, workerInvestor5)

	// print output
	counterTotal := 0
	for investor := range chanInvestorSum {
		if investor.ID > 0 {
			counterTotal++
		}
	}
	log.Printf("%d data migrated", counterTotal)

	return nil
}

func (u *Usecase) getInvestors() <-chan model.Investor {
	chanOut := make(chan model.Investor)

	go func() {
		log.Println("go routine triggerd getInvestors")
		investors, err := u.repo.GetPsql(context.Background())
		if err != nil {
			log.Printf("error get investors: %v", err.Error())
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
		log.Println("go routine triggerd upsertInvestors")
		for investor := range chanIn {
			// save to outbox table
			now := time.Now()
			payload, err := json.Marshal(investor)
			err = errors.New("paksa")
			if err != nil {
				log.Printf("error json marshal: %v", err.Error())
				return
			}
			outBox := model.Outbox{
				Identifier: investor.ID,
				Payload:    string(payload),
				Event:      "INVESTOR",
				Status:     "PENDING",
				CreatedAt:  now,
				UpdatedAt:  now,
			}
			err = u.repo.UpsertOutbox(ctx, outBox)
			if err != nil {
				log.Printf("error save to outbox %v", err.Error())
				return
			}

			// migrations
			err = u.repo.UpsertMongo(ctx, investor)
			if err != nil {
				log.Printf("error migrations: %v", err.Error())
				outBox := model.Outbox{
					Identifier: investor.ID,
					Payload:    string(payload),
					Event:      "INVESTOR",
					Status:     "FAILED",
					CreatedAt:  now,
					UpdatedAt:  now,
				}
				err = u.repo.UpsertOutbox(ctx, outBox)
				if err != nil {
					log.Printf("error save to outbox %v", err.Error())
					return
				}
				return
			}

			// delete outbox
			err = u.repo.DeleteOutbox(ctx, investor.ID)
			if err != nil {
				log.Printf("error save to outbox %v", err.Error())
				return
			}
			chanOut <- investor
		}
		close(chanOut)
	}()

	return chanOut
}

func mergeChanInvestor(chanInMany ...<-chan model.Investor) <-chan model.Investor {
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
