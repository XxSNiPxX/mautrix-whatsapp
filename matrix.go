// mautrix-whatsapp - A Matrix-WhatsApp puppeting bridge.
// Copyright (C) 2018 Tulir Asokan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	log "maunium.net/go/maulogger"
)

type MatrixListener struct {
	bridge *Bridge
	log    *log.Sublogger
	stop   chan struct{}
}

func NewMatrixListener(bridge *Bridge) *MatrixListener {
	return &MatrixListener{
		bridge: bridge,
		stop:   make(chan struct{}, 1),
		log: bridge.Log.CreateSublogger("Matrix", log.LevelDebug),
	}
}

func (ml *MatrixListener) Start() {
	for {
		select {
		case evt := <-ml.bridge.AppService.Events:
			log.Debugln("Received Matrix event:", evt)
		case <-ml.stop:
			return
		}
	}
}

func (ml *MatrixListener) Stop() {
	ml.stop <- struct{}{}
}
