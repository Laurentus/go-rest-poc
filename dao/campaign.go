package dao

import (
	"log"

	"github.com/Laurentus/poc-app/db"
	"github.com/Laurentus/poc-app/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Campaign struct {
	collection *mongo.Collection
}

// We need similar pattern for depencency management
// wire compile-time DI would autogenerate these functions for us
func NewCampaign(c *mongo.Collection) Campaign {
	return Campaign{c}
}

func Campaigns() Campaign {
	return NewCampaign(db.GetDb().Collection("campaigns"))
}

/** All returns all campaigns (that would match potential query)
 * The main value of these objects is to provide type safety
 * The types get also bound to their respective marshallers
 */
func (c Campaign) All() chan models.Campaign {
	future := make(chan models.Campaign)
	go c.findAllCampaigns(future)

	return future
}

func (c Campaign) findAllCampaigns(out chan models.Campaign) {
	campaignHandler := func(cur *mongo.Cursor) {
		var camp models.Campaign
		err := cur.Decode(&camp)
		if err != nil {
			log.Fatal(err)
		}
		// It is more efficient to provide single results as they come.
		// We'll do the list building when we need it (json output)
		out <- camp
	}
	go FindAll(c.collection, campaignHandler, func() {
		close(out)
	})
}
