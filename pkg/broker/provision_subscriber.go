package broker

import (
	"encoding/json"

	"github.com/fusor/ansible-service-broker/pkg/apb"
	"github.com/fusor/ansible-service-broker/pkg/dao"
	logging "github.com/op/go-logging"
)

type ProvisionWorkSubscriber struct {
	dao       *dao.Dao
	log       *logging.Logger
	msgBuffer <-chan WorkMsg
}

func NewProvisionWorkSubscriber(dao *dao.Dao, log *logging.Logger) *ProvisionWorkSubscriber {
	return &ProvisionWorkSubscriber{dao: dao, log: log}
}

func (p *ProvisionWorkSubscriber) Subscribe(msgBuffer <-chan WorkMsg) {
	p.msgBuffer = msgBuffer

	var pmsg *ProvisionMsg
	var extCreds *apb.ExtractedCredentials
	go func() {
		p.log.Info("Listening for provision messages")
		for {
			msg := <-msgBuffer

			p.log.Debug("Processed message from buffer")
			// HACK: this seems like a hack, there's probably a better way to
			// get the data sent through instead of a string
			json.Unmarshal([]byte(msg.Render()), &pmsg)

			if pmsg.Error != "" {
				p.dao.SetState(pmsg.InstanceUUID, apb.JobState{Token: pmsg.JobToken, State: apb.StateFailed})
			} else {
				json.Unmarshal([]byte(pmsg.Msg), &extCreds)
				p.dao.SetState(pmsg.InstanceUUID, apb.JobState{Token: pmsg.JobToken, State: apb.StateSucceeded})
				p.dao.SetExtractedCredentials(pmsg.InstanceUUID, extCreds)
			}
		}
	}()
}
