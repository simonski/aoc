/*
 * -------------------
 * |                 |
 * |                 |
 * |                 |
 * |                 |
 * |        1        |
 * |                 |
 * |                 |
 * |                 |
 * |                 |
 * -------------------
 * 
 * TODO
 * - the list and grids now have a range of classes set on the - create the class definitions for each
 *
 * - investigate DnD on add/remove
 *
 * - put a delete button on the UI
 *
 */

var App = App || {};
App.Model = {

	init : function() {
		Log.d("App.Model.init()");
	},

	/**
	 * the user profile and data they can interact with
	 */
	"apps" : [],
	"devices" : [],
	"themes" : [],
	"contacts" : [],
	"photos" : [],

	displayMode : "phone",

	setDevices : function(devices) {
		App.Model.devices = devices;
	},

	getDevices : function(devices) {
		return App.Model.devices;
	},

	setProfile: function(profile) {
		App.Model.User = profile;
	},

	getProfile: function() {
		return App.Model.User;
	},

	/**
	 * returns the contact details for a given id
	 * 
	 * @param contact_id
	 */
	getContactById : function(contact_id) {
		Log.d("App.Model.getContactById(" + contact_id + ")");
		for ( var index = 0; index < App.Model.contacts.length; index++) {
			var contact = App.Model.contacts[index];
			if (contact.contact_id == contact_id) {
				return contact;
			}
		}
		return null;
	},

	getCurrentDevice : function() {
		return App.Model.device;
	},

	getCurrentDeviceLayout : function() {
		return App.Model.deviceLayout;
	},

	/**
	 * from View.onPhoneButton Server.getDeviceConfig onSuccess
	 * 
	 * it will eventually call 'redrawDevice'
	 * 
	 */
	setCurrentDevice : function(device) {
		Log.d("App.Model.setCurrentDevice(" + device.device_id + ")");
		App.Model.device = device;
		App.View.redrawDevice();
		App.View.selectTabContacts();
	},

	setCurrentDeviceLayout : function(deviceLayout) {
		Log.d("App.Model.setCurrentDeviceLayout(" + deviceLayout.device_id + ")");
		App.Model.deviceLayout = deviceLayout;
		// App.View.redrawDeviceLayout();
		// App.View.selectTabContacts();
	},

	setCurrentDeviceByName : function(name) {
		var devices = App.Model.devices;
		var device;
		$(devices).each(function(index, dev) {
			if (dev.device_name == name) {
				device = dev;
			}
		});
		if (device) {
			App.Model.setCurrentDevice(device);
		}

	},

	foo : function(device) {
		var buttons = $(".phone_button");
		var button_to_click;
		buttons.each(function(index, btn) {
			btn = $(btn);
			if (btn.attr("device_id") == device.device_id) {
				button_to_click = btn;
			}
		});
		if (button_to_click) {
			button_to_click.click();
		}
	},

	getDisplayMode : function() {
		Log.d("App.Model.getDisplayMode()");
		return App.Model.displayMode;
	},

	setDisplayMode : function(mode) {
		Log.d("App.Model.setDisplayMode(" + mode + ")");
		App.Model.displayMode = mode;
	},

};