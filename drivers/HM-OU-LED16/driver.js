'use strict';

const Homey = require('homey');
const Driver = require('../../lib/driver.js');

class HomematicDriver extends Driver {

    onInit() {
        super.onInit();
        this.capabilities = [
            'homematic_led_status'
        ];
        this.homematicTypes = ['HM-OU-LED16'];
        this.numDevices = 16
        this.multiDevice = true

        this.log(this.homematicTypes.join(','), 'has been inited');

        this._flowTriggerButtonPressed = this.homey.flow.getDeviceTriggerCard('HM-OU-LED16-press')
            .registerRunListener((args, state) => {
                if (args.button == state.button && args.pressType == state.pressType) {
                    return Promise.resolve(true)
                } else {
                    return Promise.reject(false)
                }
            })

        this.homey.flow.getActionCard('homematic_set_led_status')
            .registerRunListener((args, state) => {
                return args.device.triggerCapabilityListener('homematic_led_status', args.led_status, {})

            })
    }

    triggerButtonPressedFlow(device, tokens, state) {
        this._flowTriggerButtonPressed
            .trigger(device, tokens, state)
            .catch(this.error)
    }
}

module.exports = HomematicDriver;
