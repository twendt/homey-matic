'use strict';

const Homey = require('homey');
const Device = require('../../lib/device.js')
const Convert = require('../../lib/convert')

const capabilityMap = {
    "onoff": {
        "channel": 3,
        "key": "STATE",
        "set": {
            "key": "STATE",
            "channel": 3
        }
    },
    "measure_power": {
        "channel": 6,
        "key": "POWER"
    },
    "measure_voltage": {
        "channel": 6,
        "key": "VOLTAGE"
    },
    "measure_current": {
        "channel": 6,
        "key": "CURRENT",
        "convert": Convert.ATomA
    },
    "meter_power": {
        "channel": 6,
        "key": "ENERGY_COUNTER",
        "convert": Convert.WhToKWh
    }
}

class HomematicDevice extends Device {

    onInit() {
        super.onInit(capabilityMap);
    }
}

module.exports = HomematicDevice;
