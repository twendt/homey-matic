'use strict';

const Homey = require('homey');
const Device = require('../../lib/device.js')
const Convert = require('../../lib/convert.js')

const capabilityMap = {
    "alarm_homematic_rain": {
        "channel": 1,
        "key": "STATE",
        "convert": Convert.toBoolean
    },
    "onoff": {
        "channel": 2,
        "key": "STATE",
        "set": {
            "key": "STATE",
            "channel": 2
        }
    }
}

class HomematicDevice extends Device {

    onInit() {
        super.onInit(capabilityMap);
    }

    initializeExtraEventListeners() {
        var self = this;
        let tokens = {};
        let state = {};
        self.bridge.on('event-' + self.deviceAddress + ':1-STATE', (value) => {
            if (value == 1) {
                self.driver.triggerTurnedOn(self, tokens, state)
            } else {
                self.driver.triggerTurnedOff(self, tokens, state)
            }
        });

    }
}

module.exports = HomematicDevice;
