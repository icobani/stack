package model

import (
	"time"

	"github.com/ilgooz/cryptoutils"
	"github.com/ilgooz/stack/conf"
	"gopkg.in/mgo.v2/bson"
)

type Token struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	UserID    bson.ObjectId `bson:"user_id"`
	Token     string
	CreatedAt time.Time `bson:"created_at"`
}

func NewToken(userID bson.ObjectId) Token {
	return Token{
		ID:        bson.NewObjectId(),
		UserID:    userID,
		Token:     cryptoutils.RandToken(conf.TokenSize),
		CreatedAt: time.Now(),
	}
}
