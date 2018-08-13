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

package config

import (
	"maunium.net/go/mautrix-appservice"
	"regexp"
)

func (config *Config) NewRegistration() (*appservice.Registration, error) {
	registration := appservice.CreateRegistration("mautrix-whatsapp")

	err := config.copyToRegistration(registration)
	if err != nil {
		return nil, err
	}

	config.AppService.ASToken = registration.AppToken
	config.AppService.HSToken = registration.ServerToken
	return registration, nil
}

func (config *Config) GetRegistration() (*appservice.Registration, error) {
	registration := appservice.CreateRegistration("mautrix-whatsapp")

	err := config.copyToRegistration(registration)
	if err != nil {
		return nil, err
	}

	registration.AppToken = config.AppService.ASToken
	registration.ServerToken = config.AppService.HSToken
	return registration, nil
}

func (config *Config) copyToRegistration(registration *appservice.Registration) error {
	registration.ID = config.AppService.ID
	registration.URL = config.AppService.Address
	registration.RateLimited = false
	registration.SenderLocalpart = config.AppService.Bot.Username

	userIDRegex, err := regexp.Compile(config.Bridge.FormatUsername("[0-9]+", "[0-9]+"))
	if err != nil {
		return err
	}
	registration.Namespaces.RegisterUserIDs(userIDRegex, true)
	return nil
}
