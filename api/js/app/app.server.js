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
App.Server = {

	// indicates if polling is currently active
	isPolling : false,
	// the interval value
	pollerId : null,

	init : function() {
		Log.d("App.Server.init()");
		/*
		 * all calls are over JSON
		 */
		$.ajaxSetup({
			headers : {
				"Accept" : "json",
	            "x-client-type": "javascript"
    		},
			cache : false
		});

	},

	getDevices : function(callback) {
		callback = Server._checkCallback(callback);
		Log.d("App.Server.getDevices");
		Server.device_list({
			onSuccess : function(data) {
				var devices = data;
				App.Model.setDevices(devices);
				// App.View.redrawDevices();
				// App.View.updateEventHandlers();
			},
			onFailure : function(err) {
				callback.onFailure(err);
			}
		});
	},

	getProfile : function(callback) {
 		callback = Server._checkCallback(callback);
		var mycb = {
			onSuccess : function(data, response) {
				var profile = _.extend(App.Objects.User, data);
				App.Model.setProfile(profile);
				callback.onSuccess(profile, response);
			},
			onFailure : function(err) {
				App.Model.setProfile(null);
				callback.onFailure(err);
			}
		};
		Server.profile(mycb);
	},

	startPolling : function() {
		if (App.Server.isPolling) {
			Log.d("App.Server.startPolling(): already polling");
			return;
		} else {
			Log.d("App.Server.startPolling()");

		}

		App.Server.isPolling = true;
		App.Server.pollerId = setInterval(App.Server.poll,
				App.Constants.POLL_PERIOD);
		amplify.publish(Server.Topic.POLLING_STARTED);
	},

	poll: function() {
//		Log.d("App.Server.poll");

		var pollState = {
			"themes_md5" : App.Model.themes.md5,
			"theme_version" : App.Constants.THEME_VERSION
		};
		if ( App.Model.deviceLayout != null ) {
			pollState.device_layout_last_modified = App.Model.deviceLayout.last_modified;
			pollState.device_id = App.Model.deviceLayout.device_id;
		};

		Server.poll(pollState, {
			onSuccess : function(data, response) {
//				Log.d("App.Server.poll() success");
				if (data.events != null) {
					for ( var i = 0; i < data.events.length; i++) {
						var event = data.events[i];
						amplify.publish(event.action, event);
					}
				}
				if (data.themes != null) {
					var themes = data.themes;
					if (App.Model.themes) {
						if (App.Model.themes.md5 !== themes.md5) {
							App.Model.themes = themes;
							App.View.redrawThemes();
						} else {
							// no change
						}
					} else {
						App.Model.themes = themes;
						App.View.redrawThemes();
					}

				}
				if ( data.device_layout != null ) {
					
//					var selectedCell = App.Model.deviceLayout.getSelectedCell();
					
					var dl = new DeviceLayout(data.device_layout);
					// find the same cell and mark it as selected
					
					var current_layout = App.Model.getCurrentDeviceLayout();
					var selected_cell = null;
					if ( current_layout != null ) {
						selected_cell = current_layout.getSelectedCell();
					}
				
//					var deviceLayout = new DeviceLayout(device_layout);
					if ( selected_cell != null ) {
						var cell = dl.getCellAt(selected_cell);
						if ( cell != null ) {
							cell.setSelected(true);
						}
					}
					
//					for ( var index=0; index<dl.cells.length; index++) {
//						var candidate = dl.cells[index];
//						if ( candidate.getId() == selectedCell.getId() && candidate.type == selectedCell.type ) {
//							candidate.setSelected(true);
//							break;
//						}
//					};
					
					App.Model.setCurrentDeviceLayout(dl);
					App.View.showDragDrop();
				}

			},
			onFailure : function(err) {
				Log.e("App.Server.poll: Server.poll failure: " + err);

			}
		});
	},

	/**
	 * stops the poller (if it is active)
	 */
	stopPolling : function() {
		if (!App.Server.isPolling) {
			Log.d("App.Server.stopPolling(): not polling");
			return;
		} else {
			Log.d("App.Server.stopPolling");
		}
		App.Server.isPolling = false;
		clearInterval(App.Server.pollerId);
		App.Server.pollerId = null;
		amplify.publish(Server.Topic.POLLING_STOPPED);
	},

	/**
	 * called by the poller when a device is registered
	 */
	onDeviceRegistered : function(event) {
		Log.d("App.Server.onDeviceRegistered - device_id is "
				+ event.data.device_id);
		App.Server.getDevices();

	},

	/**
	 * called by the poller when a device is registered
	 */
	onDeviceDeregistered : function(event) {
		Log.d("App.Server.onDeviceDeregistered - device_id is "
				+ event.data.device_id);
		App.Server.getDevices();
	},

	/**
	 * fetches up to date config from server and refreshes the UI
	 */
	getDeviceConfig : function(device_id) {
		Log.d("App.Server.getDeviceConfig(" + device_id + ")");

		Server.device_get(device_id, {
			onSuccess : function(data) {
				var device_config = data.config;
				var device_layout = data.layout;
				var device = new Device(device_config);
				var current_layout = App.Model.getCurrentDeviceLayout();
				var selected_cell = null;
				if ( current_layout != null ) {
					selected_cell = current_layout.getSelectedCell();
				}
			
				var deviceLayout = new DeviceLayout(device_layout);
				if ( selected_cell != null ) {
					var cell = deviceLayout.getCellAt(selected_cell);
					if ( cell != null ) {
						cell.setSelected(true);
					}
				}
				Log.d("Server.device_get() success");
				App.Model.setCurrentDeviceLayout(deviceLayout);
				App.Model.setCurrentDevice(device);
			},
			onFailure : function(data) {
				Log.d("App.Server.getDeviceConfig(" + data + ") failure; drawLoginPage");
				App.View.drawLoginPage();
			}
		});
	},

	saveDeviceConfig : function() {
		Log.d("App.Server.saveDeviceConfig");
		var currentDevice = App.Model.getCurrentDevice();
		if (currentDevice) {
			Server.device_put(currentDevice.device_id, currentDevice, {
				onSuccess : function() {
					Log.d("App.Server.saveDeviceConfig(): success");
				},
				onFailure : function() {
					Log.d("App.Server.saveDeviceConfig(): failure");

				}
			});
		}
	},

	deleteDevice : function() {
		Log.d("App.Server.deleteDevice");
		var currentDevice = App.Model.getCurrentDevice();
		if (currentDevice) {
			Server.deregister_device(currentDevice.device_id, {
				onSuccess : function() {

				},
				onFailure : function() {
				}
			});
		}
	},

	saveDeviceName : function(device_id, device_name) {
		Log.d("App.Server.saveDeviceName");
		Server.device_update(device_id, "device_name", device_name, {
			onSuccess : {},
			onFailure : {}
		});
	},

	saveGridColumns : function(device_id, grid_columns) {
		Log.d("App.Server.saveGridColumns");
		Server.device_update(device_id, "grid_columns", grid_columns, {
			onSuccess : {},
			onFailure : {}
		});
	},

	getSiteData : function(callback) {
		Log.d("App.Server.getSiteData");
		callback = Server._checkCallback(callback);
		Server.site_data(App.Constants.THEME_VERSION, {
			onSuccess : function(data) {
				if (data && data.devices) {
					App.Model.setDevices(data.devices);
				}
				if (data && data.contacts) {
					App.Model.contacts = data.contacts;
				}
				if (data && data.apps) {
					App.Model.apps = data.apps;
				}
				if (data && data.themes) {
					App.Model.themes = data.themes;
				}
				callback.onSuccess();
			},
			onFailure : function() {
				callback.onFailure();
			}
		});
	},

	fetchAndRedrawContacts : function() {
		Log.d("App.Server.fetchAndRedrawContacts");
		Server.contact_list({
			onSuccess : function(data) {
				Log.d("Server.contact_list: success");
				App.Model.contacts = data;
				App.View.redrawContacts();
			},
			onFailure : function() {
				Log.d("Server.contact_list: failure");
			}
		});
	}

};
