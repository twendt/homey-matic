'use strict';

const Homey = require('homey');
const Driver = require('../../lib/driver.js');

class HomematicDriver extends Driver {

    onInit() {
        super.onInit();
        this.capabilities = [
            'onoff',
            'dim'
        ];
        this.homematicTypes = ['HmIP-FDT'];
        this.log(this.homematicTypes.join(','), 'has been inited');
    }
}

module.exports = HomematicDriver;
