/**
 * array contains named html templates
 */
var Template = {};

var Server = {

    init: function() {
	    $.ajaxSetup({ "headers": { "Accept": "x-auth-json" } }); 
    },

	_checkCallback: function(url, callback) {
		if ( callback == null ) {
			callback = {
					onSuccess: function() {
                        Log.d("Server._checkCallback(" + url + "): anonymous onSuccess << YOU SHOULD IMPLEMENT SOMETHING");
					},
					onFailure: function() {
                        Log.d("Server._checkCallback(" + url + "): anonymous onFailure << YOU SHOULD IMPLEMENT SOMETHING");
   					}
			};
		}	
        if ( callback.onSuccess == null ) {
            callback.onSuccess = function() {
                Log.d("Server._checkCallback(" + url + "): anonymous onSuccess << YOU SHOULD IMPLEMENT SOMETHING");
            }
        }
        if ( callback.onFailure == null ) {
            callback.onFailure = function() {
                Log.d("Server._checkCallback(" + url + "): anonymous onFailure << YOU SHOULD IMPLEMENT SOMETHING");
            }
        }
		return callback;
	},
	
	_post: function(url, data_in, callback, successTopic, failureTopic) {
		callback = Server._checkCallback(url, callback);
		$.post(url, data_in, function(data, response) {
			try {
				var result = $.parseJSON(data);
                if ( result.status != null && result.status == "challenge" ) {
                    // not logged in
                    Log.d("Server._post('" + url + "', not logged in.");
    				callback.onFailure(result);
	    			amplify.publish(failureTopic, result);
	    			amplify.publish(Server.Topic.LOGIN_REQUIRED, result);

                } else {
					// already done
					callback.onSuccess(result);
					amplify.publish(successTopic, result);
					return;
				}	
			} catch (err) {
				callback.onFailure(err);
				amplify.publish(failureTopic, err);
			}
		});
	},

	_getCookie : function(name) {
		var cookieValue = null;
		if (document.cookie && document.cookie != '') {
			var cookies = document.cookie.split(';');
			for ( var i = 0; i < cookies.length; i++) {
				var cookie = jQuery.trim(cookies[i]);
				// Does this cookie string begin with the name we want?
				if (cookie.substring(0, name.length + 1) == (name + '=')) {
					cookieValue = decodeURIComponent(cookie
							.substring(name.length + 1));
					break;
				}
			}
		}
		return cookieValue;
	},
    
	
	_get: function(url, data_in, callback, successTopic, failureTopic) {
		callback = Server._checkCallback(url, callback);
        if ( data_in == null ) {
            data_in = {};
        }
		$.get(url, data_in, function(data, response) {
			try {
				var result = $.parseJSON(data);
                if ( result.status && result.status == "challenge" ) {
                    // not logged in
                    Log.d("Server._post('" + url + "', not logged in.");
    				callback.onFailure(result);
	    			amplify.publish(failureTopic, result);
	    			amplify.publish(Server.Topic.LOGIN_REQUIRED, result);
                } else {
    				callback.onSuccess(result);
	    			amplify.publish(successTopic, result);
                }
                    
			} catch (err) {
				callback.onFailure(err);
				amplify.publish(failureTopic, err);
			}
		});
		
	},

	Topic: {

        GET_DOCS_SUCCESS: "docs_success",
        GET_DOCS_FAILURE: "docs_failure",

        LOGIN_REQUIRED: "login_required",
		LOGIN_SUCCESS: "login_success",
		LOGIN_FAILURE: "login_failure",

		LOGOUT_SUCCESS: "logout_success",
		LOGOUT_FAILURE: "logout_failure",

   		REGISTER_SUCCESS: "register_success",
		REGISTER_FAILURE: "register_failure",

        GET_FILE_LIST_SUCCESS: "get_file_list_success",
        GET_FILE_LIST_FAILURE: "get_file_list_failure",

		GET_PROFILE_SUCCESS: "get_profile_success",
		GET_PROFILE_FAILURE: "get_profile_failure",
		TEMPLATE_DOWNLOAD_SUCCESS: "template_downloaded_success",
		TEMPLATE_DOWNLOAD_FAILURE: "template_downloaded_failure",
		DEVICE_GET_SUCCESS: "get_device_success",
		DEVICE_GET_FAILURE: "get_device_failure",
		DEVICE_PUT_SUCCESS: "put_device_success",
		DEVICE_PUT_FAILURE: "put_device_failure",
		DEVICE_UPDATE_SUCCESS: "update_device_success",
		DEVICE_UPDATE_FAILURE: "update_device_failure",
		DEVICE_LIST_SUCCESS: "device_list_success",
		DEVICE_LIST_FAILURE: "device_list_failure",
		SITE_DATA_SUCCESS: "site_data_success",
		SITE_DATA_FAILURE: "site_data_failure",

		ADD_CONTACT_SUCCESS: "add_contact_success",
		ADD_CONTACT_FAILURE: "add_contact_failure",

		DELETE_CONTACT_SUCCESS: "delete_contact_success",
		DELETE_CONTACT_FAILURE: "delete_contact_failure",

		UPDATE_CONTACT_SUCCESS: "update_contact_success",
		UPDATE_CONTACT_FAILURE: "update_contact_failure",

		LIST_CONTACT_SUCCESS: "list_contact_success",
		LIST_CONTACT_FAILURE: "list_contact_failure",

		POLL_SUCCESS: "poll_success",
		POLL_FAILURE: "poll_failure",

		POLL_THEME_SUCCESS: "poll_theme_success",
		POLL_THEME_FAILURE: "poll_theme_failure",

		// called by the poller when an event occurs
		POLLER_DEVICE_REGISTERED: "device_register",
		
		POLLER_DEVICE_DEREGISTERED: "device_deregister",

		POLLER_COMMAND: "poller_command",

	},

	login: function(username, password, callback) {
		
        var url = "/api/user/login";
        callback = Server._checkCallback(url, callback);
		$.post(url, { "username": username, "x-auth-json": "yes" }, function(data) {
			var result_1 = $.parseJSON(data);
			if ( result_1.status == "success" ) {
				// already done
				callback.onSuccess(result_1);
				amplify.publish(Server.Topic.LOGIN_SUCCESS, result_1);
				return;
			}
			
			if ( result_1.status == "challenge" ) {
				// we have the profile
                //

                var password_md5 = MD5(password);

                var ha1 = MD5(username + ":" + result_1.realm + ":" + password_md5);
                var ha2 = MD5("POST:" + result_1.uri);
                var ha3 = [ha1, 
                           result_1.nonce,
                           result_1.nc,
                           result_1.cnonce,
                           result_1.qop,
                           ha2].join(":");
                var response = MD5(ha3);

                var pwd = response;
                
                var pwd = MD5(password_md5 + result_1.nonce);
				
				$.post("/api/user/login", 
                    { 
                    "username": username, 
                    "response": pwd, 
                    "nc": result_1.nc,
                    "nonce": result_1.nonce,
                    "cnonce": result_1.cnonce,
                    "qop": result_1.qop,
                    "uri": result_1.uri,
                    "salt": result_1.salt 
                    }
                    
                    , function(data) {
					var result_2 = $.parseJSON(data);
					if ( result_2.status == "success" ) {
						// we have logged in - fetch the profile
                        Log.d("Server.login: 2nd challenge succcess, logged in.");
						callback.onSuccess(result_2);
						amplify.publish(Server.Topic.LOGIN_SUCCESS, result_2);
					} else { // if ( result_2.status == "failure" ) {
                        Log.d("Server.login: 2nd challenge failed, could be password failure.");
						callback.onFailure(result_2);
						amplify.publish(Server.Topic.LOGIN_FAILURE, result_2);
					}
				});
				
			} else {
				// we do not
				callback.onFailure(result_1);
				amplify.publish(Server.Topic.LOGIN_FAILURE, result_1);
			}
		});
	},
			
	register: function(username, email, password, callback) {

        var url = "/api/user/register";
        Server._checkCallback(url, callback);
		
		$.post(url, { "username": username, "email": email, "password": password }, function(data) {
			var result = $.parseJSON(data);
			if ( result.status == "success" ) {
				// already done
				amplify.publish(Server.Topic.REGISTER_SUCCESS, result);
				return;
			} else {
				amplify.publish(Server.Topic.REGISTER_FAILURE, result);
				return;
				
			}
		});
		
	},
	
	
	logout: function() {
		$.get("/api/user/logout", null, function(data) {
			var result = $.parseJSON(data);
			if ( result.status == "success" ) {
				Template = {};
                Log.d("Server.logout() Logged out");
				amplify.publish(Server.Topic.LOGOUT_SUCCESS, result);
			} else {
				amplify.publish(Server.Topic.LOGOUT_FAILURE, result);
			}
		});
	},
			
	profile: function(username, callback) {
		Server._get("/api/user/profile/search", {"username": username}, callback, Server.Topic.GET_PROFILE_SUCCESS, Server.Topic.GET_PROFILE_FAILURE);
	},


	/**
	 * fetches the template from the server and stores it in Template
	 * @param templateName
	 * @param callback
	 */
	template: function(templateName, callback) {
		if ( callback == null ) {
			callback = {
				onSuccess: function() {
				},
				onFailure: function() {
				}
			};
		}
		$.post("/api/template", { "template": templateName }, function(data) {
			Template[templateName] = data;
			callback.onSuccess();
			amplify.publish(Server.Topic.TEMPLATE_DOWNLOAD_SUCCESS, templateName);
		}); 
	},
		

	file_list: function(callback) {
		Server._get("/api/file/list", null, callback, Server.Topic.GET_FILE_LIST_SUCCESS, Server.Topic.GET_FILE_LIST_FAILURE);
	},
        	
	mongo_collections: function(command) {
		var data = { "action": command };
		Server._post("/api/admin/mongo", data)

	},

	mongo_get_data: function(name) {
		var data = { "action": "get", "name": name };
		Server._post("/api/admin/mongo", data)
	},

	poll: function(data_in, callback) {
		Server._post("/api/site_poll", data_in, callback, Server.Topic.POLL_SUCCESS, Server.Topic.POLL_FAILURE);
	},

	 
		
	docs: function(callback) {
		Server._get("/api/docs", null, callback, Server.Topic.GET_DOCS_SUCCESS, Server.Topic.GET_DOCS_FAILURE);
	},
    
	device_get: function(device_id, callback) {
		var data = { "device_id": device_id };
		Server._post("/api/device_get", data, callback, Server.Topic.DEVICE_GET_SUCCESS, Server.Topic.DEVICE_GET_FAILURE);
	},
	
	layout_get: function(device_id, callback) {
		var data = { "device_id": device_id };
		Server._get("/api/device_layout", data, callback, Server.Topic.DEVICE_LAYOUT_GET_SUCCESS, Server.Topic.DEVICE_LAYOUT_GET_FAILURE);
	},

	layout_put: function(device_id, data, callback) {
		var json_text = JSON.stringify(data); // , null, 2);
		var data_in = { "device_id": device_id, "json": json_text };
		Server._post("/api/device_layout", data_in, callback, Server.Topic.DEVICE_LAYOUT_PUT_SUCCESS, Server.Topic.DEVICE_LAYOUT_PUT_FAILURE);
	},

	device_put: function(device_id, data, callback) {
		var json_text = JSON.stringify(data); // , null, 2);
		var data_in = { "device_id": device_id, "json": json_text };
		Server._post("/api/device_put", data_in, callback, Server.Topic.DEVICE_PUT_SUCCESS, Server.Topic.DEVICE_PUT_FAILURE);
	},

	device_update: function(device_id, key, value, callback) {
		var data_in = { "device_id": device_id, "key": key, "value": value };
		$.post("/api/device_update", data_in, callback, Server.Topic.DEVICE_UPDATE_SUCCESS, Server.Topic.DEVICE_UPDATE_FAILURE);
	},

	device_update_setting: function(device_id, key, value, callback) {
		var data_in = { "device_id": device_id, "key": key, "value": value };
		$.post("/api/device_update_setting", data_in, callback, Server.Topic.DEVICE_SETTING_UPDATE_SUCCESS, Server.Topic.DEVICE_SETTING_UPDATE_FAILURE);
	},

	device_add_contact: function(device_id, contact_id, callback) {
		var data_in = { "device_id": device_id, "contact_id": contact_id };
		$.post("/api/device_add_contact", data_in, callback, Server.Topic.DEVICE_ADD_CONTACT_SUCCESS, Server.Topic.DEVICE_ADD_CONTACT_FAILURE);
	},

	device_remove_contact: function(device_id, contact_id, callback) {
		var data_in = { "device_id": device_id, "contact_id": contact_id };
		$.post("/api/device_remove_contact", data_in, callback, Server.Topic.DEVICE_RM_CONTACT_SUCCESS, Server.Topic.DEVICE_RM_CONTACT_FAILURE);
	},

	device_add_app: function(device_id, app_id, callback) {
		var data_in = { "device_id": device_id, "app_id": app_id };
		$.post("/api/device_add_app", data_in, callback, Server.Topic.DEVICE_ADD_CONTACT_SUCCESS, Server.Topic.DEVICE_ADD_CONTACT_FAILURE);
	},

	device_remove_app: function(device_id, app_id, callback) {
		var data_in = { "device_id": device_id, "app_id": app_id };
		$.post("/api/device_remove_app", data_in, callback, Server.Topic.DEVICE_RM_APP_SUCCESS, Server.Topic.DEVICE_RM_APP_FAILURE);
	},

	device_move_app: function(device_id, page_id, app_id, row, position, callback) {
		var data_in = { "device_id": device_id, "page_id": page_id, "app_id": app_id, "row": row, "position": position };
		$.post("/api/device_move_app", data_in, callback, Server.Topic.DEVICE_MV_APP_SUCCESS, Server.Topic.DEVICE_MV_APP_FAILURE);
	},

	device_move_cells: function(device_id, cells, callback) {
		var json_text = JSON.stringify(cells); // , null, 2);
		var data_in = { "device_id": device_id, "json": json_text };
		$.post("/api/device_move_cells", data_in, callback, Server.Topic.DEVICE_LAYOUT_PUT_SUCCESS, Server.Topic.DEVICE_LAYOUT_PUT_FAILURE);
	},
	
	device_hide_app: function(device_id, page_id, app_id, callback) {
		var data_in = { "device_id": device_id, "page_id": page_id, "app_id": app_id};
		$.post("/api/device_hide_app", data_in, callback, Server.Topic.DEVICE_HIDE_APP_SUCCESS, Server.Topic.DEVICE_HIDE_APP_FAILURE);
	},

	device_hide_contact: function(device_id, page_id, contact_id, callback) {
		var data_in = { "device_id": device_id, "page_id": page_id, "contact_id": contact_id};
		$.post("/api/device_hide_contact", data_in, callback, Server.Topic.DEVICE_HIDE_CONTACT_SUCCESS, Server.Topic.DEVICE_HIDE_CONTACT_FAILURE);
	},

	device_move_contact: function(device_id, page_id, contact_id, row, position, callback) {
		var data_in = { "device_id": device_id, "page_id": page_id, "contact_id": contact_id, "row": row, "position": position };
		$.post("/api/device_move_contact", data_in, callback, Server.Topic.DEVICE_MV_CONTACT_SUCCESS, Server.Topic.DEVICE_MV_CONTACT_FAILURE);
	},
	
	device_list: function(callback) {
		Server._get("/api/device_list", null, callback, Server.Topic.DEVICE_LIST_SUCCESS, Server.Topic.DEVICE_LIST_FAILURE);
	},

	deregister_device: function(device_id, callback) {
		var data_in = { "device_id": device_id, };
		Server._post("/api/deregister_device", data_in, callback, Server.Topic.DEREGISTER_DEVICE_SUCCESS, Server.Topic.DEREGISTER_DEVICE_FAILURE);
	},
	
	contact_add: function(contact_name, contact_number, contact_email, contact_class, callback) {
		var data_in = { "contact_name": contact_name, 
					"contact_number": contact_number, 
					"contact_email": contact_email,
					"contact_class": contact_class
					};
		Server._post("/api/contact_create", data_in, callback, Server.Topic.ADD_CONTACT_SUCCESS, Server.Topic.ADD_CONTACT_FAILURE);
	},

	contact_list: function(callback) {
		Server._get("/api/contact_list", null, callback, Server.Topic.LIST_CONTACT_SUCCESS, Server.Topic.LIST_CONTACT_FAILURE);
	},

	contact_delete: function(contact_id, callback) {
		var data_in = { "contact_id": contact_id };
		Server._post("/api/contact_delete", data_in, callback, Server.Topic.DELETE_CONTACT_SUCCESS, Server.Topic.DELETE_CONTACT_FAILURE);
	},

	contact_update: function(contact_id, contact_name, contact_number, contact_email, callback) {
		var data_in = { "contact_id": contact_id, "contact_name": contact_name, "contact_number": contact_number, "contact_email": contact_email };
		Server._post("/api/contact_update", data_in, callback, Server.Topic.UPDATE_CONTACT_SUCCESS, Server.Topic.UPDATE_CONTACT_FAILURE);
	},

	site_data: function(theme_version, callback) {
		Server._post("/api/site_data", { "theme_version": theme_version }, callback, Server.Topic.SITE_DATA_SUCCESS, Server.Topic_SITE_DATA_FAILURE);
	},

	theme_list: function(callback) {
		Server._post("/api/theme_list", null, callback, Server.Topic.THEME_LIST_SUCCESS, Server.Topic_THEME_LIST_FAILURE);
	},

	theme_manifest: function(theme_id, callback) {
		var data = { "theme_id": theme_id };
		Server._post("/api/theme_manifest", data, callback, Server.Topic.THEME_MANIFEST_SUCCESS, Server.Topic_THEME_MANIFEST_FAILURE);
	},

};
