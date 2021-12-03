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

App.View = {

	init : function() {
		Log.d("App.View.init()");

		if (App.Constants.IS_IPAD) {
			// inject the ipad css
			Common.injectCSS("/site/css/app/ipad.css");
		} else {
			// inject the browser css
			Common.injectCSS("/site/css/app/browser.css");
		}

		// $(".btn_delete_device").click(App.View.onDeleteDeviceClicked);
		// $(".btn_reset_device_layout").click(App.View.onResetDeviceLayoutClicked);

		amplify.subscribe(Server.Topic.LOGOUT_SUCCESS, App.View.drawLoginPage);
		amplify.subscribe(Server.Topic.LOGOUT_FAILURE, App.View.drawLoginPage);

		amplify.subscribe(Server.Topic.POLLER_DEVICE_REGISTERED, App.Server.onDeviceRegistered);
		amplify.subscribe(Server.Topic.POLLER_DEVICE_DEREGISTERED, App.Server.onDeviceDeregistered);

		amplify.subscribe(Server.Topic.POLLER_COMMAND, function(data) {
			if (data.command == "logout") {
				Server.logout();
			} else if (data.command == "reload") {
				window.location.reload(true);
			}
		});

		window.onkeypress = App.View.onKeyPress;
		
	},

	/**
	 * shows the contacts tab
	 */
	selectTabContacts : function() {
		$("#tabContacts").tab().show();
		$("#tabApps").tab().hide();
		$("#tabThemes").tab().hide();
		$("#tabPhotos").tab().hide();
		$("#tabSettings").tab().hide();
		$("#tabAdmin").tab().hide();
		$(".tabHeader").removeClass("active");
		$("#tabHeaderContacts").addClass("active");
	},
	/**
	 * shows the apps tab
	 */
	selectTabApps : function() {
		$("#tabContacts").tab().hide();
		$("#tabApps").tab().show();
		$("#tabThemes").tab().hide();
		$("#tabPhotos").tab().hide();
		$("#tabSettings").tab().hide();
		$("#tabAdmin").tab().hide();
		$(".tabHeader").removeClass("active");
		$("#tabHeaderApps").addClass("active");
	},
	/**
	 * shows the themes tab
	 */
	selectTabThemes : function() {
		$("#tabContacts").tab().hide();
		$("#tabApps").tab().hide();
		$("#tabThemes").tab().show();
		$("#tabPhotos").tab().hide();
		$("#tabSettings").tab().hide();
		$("#tabAdmin").tab().hide();
		$(".tabHeader").removeClass("active");
		$("#tabHeaderThemes").addClass("active");
	},
	selectTabSettings : function() {
		$("#tabContacts").tab().hide();
		$("#tabApps").tab().hide();
		$("#tabThemes").tab().hide();
		$("#tabPhotos").tab().hide();
		$("#tabSettings").tab().show();
		$("#tabAdmin").tab().hide();
		$(".tabHeader").removeClass("active");
		$("#tabHeaderSettings").addClass("active");
	},
	selectTabAdmin : function() {
		$("#tabContacts").tab().hide();
		$("#tabApps").tab().hide();
		$("#tabThemes").tab().hide();
		$("#tabPhotos").tab().hide();
		$("#tabSettings").tab().hide();
		$("#tabAdmin").tab().show();
		$(".tabHeader").removeClass("active");
		$("#tabHeaderAdmin").addClass("active");
	},

	onActivityStarted : function() {
		$(".activity").show();
	},

	onActivityEnded : function() {
		$(".activity").hide();
	},

	clear : function() {
		Log.d("App.View.clear");
		$("#root").html("");
	},

	onResetDeviceLayoutClicked : function() {
		Log.d("App.View.onResetDeviceLayoutClicked");
		var modalData = {
			header : "Reset device?",
			body : "Do you want to reset (clear) this device?",
			yesButton : true,
			noButton : true,
			onYes : function() {
				var deviceLayout = App.Model.getCurrentDeviceLayout();
				deviceLayout.reset();
				deviceLayout.save();
				App.View.showDragDrop();
			},
			onNo : function() {
				
			}
		};
		Common.showModal(modalData);

	},

	onDeleteDeviceClicked : function() {
		Log.d("App.View.onDeleteDeviceClicked");
		var modalData = {
			header : "Delete device?",
			body : "Do you want to delete this device?",
			yesButton : true,
			noButton : true,
			onYes : function() {
				App.Server.deleteDevice();
			},
			onNo : function() {
			}
		};
		Common.showModal(modalData);
	},

	setCurrentTheme : function(theme_id) {
		// get the theme node
		var theme_node = $("#theme_holder");
		var current_theme_id = theme_node.attr("theme_id");
		if (current_theme_id != theme_id) {
			theme_node.attr("theme_id", theme_id);
			theme_node.html("");

			var theme_url = App.Constants.THEME_URL + theme_id
					+ "/theme_web.css";
			var theme_html = "<link rel='stylesheet' type='text/css' href='"
					+ theme_url + "'/>";

			theme_node.html(theme_html);
		}
	},

	/**
	 * draws the login dialog
	 */
	drawLoginPage : function() {
		Log.d("App.View.drawLoginPage()");
		App.Server.stopPolling();
		App.View.clear();
		var node = $("#root");
        node.html("");

		var html = _.template($("#template_login").html());
		node.append(html({
			"appname" : "sq1"
		}));
		var loginButton = $("#button_login");

		var usernameNode = $("#username");
		var passwordNode = $("#password");

		Common.onEnterPressed(usernameNode, function() {
			passwordNode.focus();
		});
		Common.onEnterPressed(passwordNode, function() {
			loginButton.focus();
		});

		loginButton.click(function() {
            App.View.onLoginButtonClicked();
        });

   		usernameNode.focus();


	},

	onLoginButtonClicked : function() {
		var usernameNode = $("#username");
		var passwordNode = $("#password");
		var username = usernameNode.val();
		var password = passwordNode.val();

		Server.login(username, password,
			{
			onSuccess : function(data) {
				App.Server.getProfile({
                    onSuccess : function() {
						App.Server.getSiteData({
                            onSuccess : function() {
								Log.d("App.View.onLoginButtonClicked: user is logged in, show the normal page");
								App.View.drawLoggedInPage();
							},
						});
					},
					onFailure : function(err) {
						Log	.d("App.View.onLoginButtonClicked: user is not logged in; show the login page");
						App.View.drawLoginPage();
					}
				});
			},
			onFailure : function(data) {
                // should be taken care of by the redraw
			}
		});

	},

	/**
	 * once logged in, the main page will be displayed
	 */
	drawLoggedInPage : function() {

		Log.d("App.View.drawLoggedInPage");
		App.View.clear();

		if ( App.Constants.IS_BROWSER ) {
			Common.injectCSS("/site/css/app/browser.css");
		} else if ( App.Contants.IS_IPAD ) {
			Common.injectCSS("/site/css/app/ipad.css");
		}

		var node = $("#root");
        node.html("");
		var html = _.template($("#template_loggedin").html());
		node.append(html({
			"appname" : "sq1",
			"username" : App.Model.User.getName()
		}));

		$("#username").html(App.Model.User.getName());

		$(".button_navbar").unbind().click(function(clicked_elt) {
			$(".button_navbar").each(function(index, elt) {
				$(elt).parent().removeClass("active");
			});
			$(this).parent().addClass("active");
			var template_name = $(this).attr("template_name");
			console.log("button_navbar.click(); template_name=" + template_name);
			var template_container = $("#template_container");
			template_container.html(template_name);
			App.View.drawTemplate(template_name, template_container);
		});

		// $("#button_home").click(App.View.onHomeButtonClicked);
		// $("#button_profile").click(App.View.onProfileButtonClicked);
		// $("#button_settings").click(App.View.onSettingsButtonClicked);
		// $("#button_logs").click(App.View.onLogsButtonClicked);
		// $("#button_users").click(App.View.onUsersButtonClicked);
		// $("#button_devices").click(App.View.onDevicesButtonClicked);
		// $("#button_groups").click(App.View.onGroupButtonClicked);

		if ( !App.Model.User.isAdmin() ) {
			$("#button_admin").detach();
			$("#button_users").detach();
			$("#button_groups").detach();
			$("#button_devices").detach();
			$("#button_settings").detach();
			$("#button_logs").detach();
		}
		$("#button_logout").unbind().click(App.View.onLogoutButtonClicked);

		App.View.redrawScreen();
 		$("#button_home").click();

		// App.Server.startPolling();

	},

	onLogoutButtonClicked : function() {
		Log.d("App.View.onLogoutButtonClicked");
		Server.logout();
	},

	// onOverviewButtonClicked : function() {
	// 	Log.d("App.View.onOverviewButtonClicked");
	// 	$(".button_menu").removeClass("active");
	// 	$(".button_overview").addClass("active");
	// 	App.View.drawOverview();
	// },

	// onMyProfileButtonClicked : function() {
	// 	Log.d("App.View.onMyProfileButtonClicked");
	// 	$(".button_menu").removeClass("active");
	// 	$(".button_my_profile").addClass("active");
	// 	App.View.drawProfile();
	// },

	/**
	 * retrieves from the tempalte store or fetches the template from the server
	 * 
	 * @param templateName the name of the template to draw
	 * @param node - the element to draw into
	 * @param callback
	 */
	drawTemplateFromServer : function(templateName, node, callback) {
		Log.d("App.View.drawTemplateFromServer(" + templateName + ")");
		callback = Server._checkCallback(callback);

		if (Template[templateName] != null) {
			var templateHtml = Template[templateName];
			node.html(templateHtml);
			callback.onSuccess();
		} else {
			Server.template(templateName, {
				onSuccess : function() {
					var templateHtml = Template[templateName];
					node.html(templateHtml);
					callback.onSuccess();
				},
				onFailure : function() {
					Log.d("template didn't download");
					callback.onFailure();
				}
			});
		}

	},

	drawTemplate: function(template_name, node, templateData) {
		Log.d("App.View.drawTemplate(" + template_name + ")");
		var html = _.template($("#" + template_name).html());
		if ( !templateData ) {
			templateData = {};
		}
		node.html(html(templateData));
	},

	// as drawTemplate, but displays as a modal
	drawTemplateAsDialog : function(templateName, node, callback) {

		var html = _.template($("#" + templateName).html());
		node.append(html(templateData));

		Log.d("App.View.drawTemplateAsDialog(" + templateName + ")");
		callback = Server._checkCallback(callback);

		Server.template(templateName, {
			onSuccess : function() {
				var templateHtml = Template[templateName];
				var modalNode = $(templateHtml);
				modalNode.unbind();
				node.append(modalNode);
				modalNode.modal();
				callback.onSuccess();
			},
			onFailure : function() {
				Log.d("template didn't download");
				callback.onFailure();
			}
		});

	},

	// completely redraws the screen - don't call this too often. It'll rebuild
	// the UI based on the data in the App.data
	redrawScreen : function() {
		Log.d("App.View.redrawScreen");
		try {
			App.View.redrawDevices();
			App.View.redrawContacts();
			App.View.redrawApps();
			App.View.redrawThemes();
			App.View.updateEventHandlers();

			var device = App.Model.getCurrentDevice();
			if (device != null
					|| (App.Model.devices && App.Model.devices.length > 0)) {
				var device = App.Model.devices[0];
				var device_id = device.device_id;
				// the act of fetching the device will update the screen
				App.Server.getDeviceConfig(device_id);
			}
		} catch (err) {
			alert("Redraw screen: " + err);
		}

	},

	/**
	 * reattaches all events
	 */
	updateEventHandlers : function() {
		Log.d("App.View.updateEventHandlers");

		var defaultClickFn = function() {
			alert("Clicked");
		};

		$("#tabHeaderContacts").unbind();
		$("#tabHeaderContacts").click(App.View.selectTabContacts);
		$("#tabHeaderApps").unbind();
		$("#tabHeaderApps").click(App.View.selectTabApps);
		$("#tabHeaderThemes").unbind();
		$("#tabHeaderThemes").click(App.View.selectTabThemes);
		$("#tabHeaderPhotos").unbind();
		$("#tabHeaderPhotos").click(App.View.selectTabPhotos);
		$("#tabHeaderSettings").unbind();
		$("#tabHeaderSettings").click(App.View.selectTabSettings);
		$("#tabHeaderAdmin").unbind();
		$("#tabHeaderAdmin").click(App.View.selectTabAdmin);

		/*
		 * $("#button_map").click(App.View.onMapButtonClicked);
		 * $("#button_grid_list").unbind();
		 * $("#button_grid_list").click(App.View.onToggleGridListButtonClicked);
		 */

		/*
		 * $(".button_grid").unbind();
		 * $(".button_grid").click(App.View.onGridButtonClicked)
		 */

		$(".button_list").unbind();
		$(".button_list").click(App.View.onListButtonClicked);

		$(".button_dragdrop").unbind();
		$(".button_dragdrop").click(App.View.onDragDropButtonClicked);

		$(".button_grid_size").unbind();
		$(".button_grid_size").click(App.View.onGridSizeButtonClicked);

		$("#button_add_an_app").click(defaultClickFn);
		$("#button_add_a_contact").click(defaultClickFn);

		$(".btn_delete_device").click(App.View.onDeleteDeviceClicked);
		$(".btn_reset_device_layout")
				.click(App.View.onResetDeviceLayoutClicked);

	},

	onDeleteDeviceClicked : function() {
		var modalData = {
			header : "Delete device?",
			body : "Do you want to delete this device?",
			yesButton : true,
			noButton : true,
			onYes : function() {
				App.Server.deleteDevice();
				App.Server.getSiteData({
					onSuccess : function() {
						Log.d("user is logged in, show the normal page");
						App.View.drawLoggedInPage();
					},
					onFailure : function() {
						Log.d("user is not logged in; show the login page");
						App.View.drawLoginPage();
					}
				});
			},
			onNo : function() {

			}
		};
		Common.showModal(modalData);
	},

	redrawApps : function() {
		Log.d("App.View.redrawApps");

		var container = $("#apps_container");
		container.html("");

		var currentDevice = App.Model.getCurrentDevice();
		if (currentDevice == null) {
			return;
		}

		var theme_id = currentDevice.theme_id;

		for ( var i = 0; i < App.Model.apps.length; i++) {
			var app = App.Model.apps[i];
			var variables = {
				packageName : app.packageName,
				label : app.label,
				img_src : App.Constants.THEME_URL + theme_id + "/apps/"
						+ app.packageName + ".png",
			};

			var html = Common.template("#template_app_button", variables);
			container.append(html);
		}

		$(".app_button").click(App.View.onAppButtonClicked);

	},

	redrawThemes : function() {
		Log.d("App.View.redrawThemes()");
		var container = $("#themes_container");
		container.html("");
		if (!App.Model.themes)
			return;
		if (!App.Model.themes.themes)
			return;
		var themes = App.Model.themes.themes;
		for ( var i = 0; i < themes.length; i++) {
			var theme = themes[i];
			var variables = {
				theme_id : theme.theme_id,
				theme_name : theme.theme_name,
				img_src : App.Constants.THEME_URL + theme.theme_id
						+ "/images/background.png",
			};

			var html = Common.template("#template_theme_button", variables);
			container.append(html);

		}

		$(".theme_button").click(App.View.onThemeButtonClicked);

	},

	/**
	 * draws the list of contacts
	 */
	redrawContacts : function() {
		Log.d("App.View.redrawContacts");

		var container = $("#contacts_container");
		container.html("");

		var currentDevice = App.Model.getCurrentDevice();
		if (currentDevice == null) {
			return;
		}
		var theme_id = currentDevice.theme_id;

		for ( var i = 0; i < App.Model.contacts.length; i++) {

			var contact = App.Model.contacts[i];
			var contact_id = contact.contact_id;
			var contact_class = contact.contact_class;

			var editor_id = "editor_" + contact_id;

			var variables = {
				contact_name : contact.contact_name,
				contact_id : contact.contact_id,
				img_src : App.Constants.THEME_URL + theme_id + "/contacts/"
						+ contact_class + ".png",
			};

			var html = Common.template("#template_contact_button", variables);
			container.append(html);

			var entry = html;

			// container.append(entry);

//    		entry.addClass("ppe");
//			entry.popover({
//			    html : true,
//			    placement: "top",
//			    title: contact_id,
//			    trigger: "click",
//			    content: function() {
//			    	var result = "<button class='btn btn-block'><i class='icon-ok'></i> Add to Phone</button>"
//			    			   + "<button class='btn btn-block'><i class='icon-wrench'></i> Edit</button>"
//			    			   + "<button class='btn btn-block btn-danger'><i class='icon-remove'></i> Delete</button>";
//			    	return result;
//			        // return $('#popover-content').html();
//			    }
//			});
			// $(".contact_button").click(App.View.onContactButtonClicked);

			var deleteBtn = entry.find(".person-editor");
			deleteBtn
					.click(function(evt) {
						var btn = $(this);
						var contact_id = btn.attr("contact_id");
						var contact_name = btn.attr("contact_name");
						var options = {
							header : "Delete Contact?",
							body : "This will permanently delete '"
									+ contact_name
									+ "' from the address book and all devices.  Do you really want to?",
							yesButton : true,
							noButton : true,
							onYes : function() {
								Server
										.contact_delete(
												contact_id,
												{
													onSuccess : function() {

														App.Model.getCurrentDevice().removeContact(contact_id);
														App.Model.getCurrentDeviceLayout().removeContact(contact_id);
														App.View.redrawDevice();
														App.Model.getCurrentDevice().save();
														App.Model.getCurrentDeviceLayout().save();

														Log.d("Server.contact_delete: success");
														App.Server.fetchAndRedrawContacts();
													},
													onFailure : function(e) {
														Log.d("Server.contact_delete: failure " + e);

													}
												});
							},
							onNo : function() {

							}
						};
						Common.showModal(options);
					});

		}

		$(".contact_button").click(App.View.onContactButtonClicked);

		var entry = $("<li class=\"scrollable_list_entry contact_button\">"
				+ "<a id=\"button_add\" href=\"#tabAdd\" data-toggle=\"tab\">"
				+ "<span class=\"person\">"
				+ "<img class=\"person-image image-up\" width=\"80px\" height=\"80px\" color=\"black\">"
				+ "<span class=\"person-text\">Add</span>" + "</span>" + "</a>"
				+ "</li>");
		container.append(entry);
		entry.click(App.View.drawAddContactPage);

	},

	/**
	 * redraws the list of devices at the top of the screen
	 */
	redrawDevices : function() {
		Log.d("App.View.redrawDevices");

		var device_container = $("#device_container");
		device_container.html("");

		var contact_images = [ "/img/contacts/grandad.png",
				"/img/contacts/billy.png", "/img/contacts/daddy.png",
				"/img/contacts/mummy.png", "/img/contacts/granny.png",
				"/img/contacts/grandad.png", "/img/contacts/billy.png",
				"/img/contacts/daddy.png", "/img/contacts/mummy.png",
				"/img/contacts/granny.png", "/img/contacts/grandad.png",
				"/img/contacts/billy.png", "/img/contacts/daddy.png",
				"/img/contacts/mummy.png", "/img/contacts/granny.png",

		];

		for ( var i = 0; i < App.Model.devices.length; i++) {
			var device = App.Model.devices[i];
			var device_id = device.device_id;

			var src = contact_images[i];

			var variables = {
				device_id : device.device_id,
				img_src : src,
				device_name : device.device_name
			};
			var html = Common.template("#template_phone_button", variables);

			device_container.append(html);
		}

		$(".phone_button").click(App.View.onPhoneButtonClicked);

	},

	redrawDevice : function() {
		Log.d("App.View.redrawDevice");

		var device = App.Model.getCurrentDevice();
		var btns = $(".phone_button");

		try {
			btns.each(function(index, btn) {
				btn = $(btn);
				if (btn.attr("device_id") == device.device_id) {
					btn.addClass("active");

					btn.find(".person_background_image").addClass("active");
					btn.find(".person_text").addClass("active");

				} else {
					btn.removeClass("active");
					btn.find(".person_background_image").removeClass("active");
					btn.find(".person_text").removeClass("active");
				}

			});
		} catch (e) {
			alert(e);
		}

		$("#content_phone").show();

		$("#button_grid").removeClass("active");
		$("#button_gridlist").removeClass("active");

		$("#content_phone").removeClass("theme_starwars");
		$("#content_phone").removeClass("theme_basic");
		$("#content_phone").removeClass("theme_tinkerbell");
		$("#content_phone").addClass("theme_" + device.theme_id);

		if (App.Model.getDisplayMode() == "phone") {
			$(".subcontent_phone_container").show();
			if (App.Model.getCurrentDevice().displayMode == "grid") {
				App.View.showGrid();
			} else if (App.Model.getCurrentDevice().displayMode == "dragdrop") {
				App.View.showDragDrop();
			} else {
				App.View.showList();
			}
		} else {
			$(".subcontent_map").show();
		}

		var device_name_node = $("#device_name");
		device_name_node.val(device.device_name);
		device_name_node.unbind().change(function() {
			var device_id = device.device_id;
			var val = device_name_node.val();// + " (" +
			// App.Model.getCurrentDevice().username
			// + ")";
			$(".person-text_" + device_id).html(val);
			device.device_name = val;
			App.Server.saveDeviceName(device_id, val);
		});

		$(".grid_columns_text").html("Columns (" + device.grid_columns + ")");
		$("#grid_columns").unbind().change(function() {
			var device_id = device.device_id;
			var val = $("#grid_columns").val();
			if (val != null) {
				device.grid_columns = val;
				$(".grid_columns_text").html("Columns (" + val + ")");
				App.Server.saveGridColumns(device_id, val);
				App.View.redrawDevice();
			}

		});

		var entries = [];
		if (device.contacts) {
			for ( var i = 0; i < device.contacts.length; i++) {
				var contact = device.contacts[i];
				if (contact) {
					entries.push({
						"type" : "contact",
						data : contact
					});
				}
			}
		}
		if (device.apps) {
			for ( var i = 0; i < device.apps.length; i++) {
				var app = device.apps[i];
				if (app) {
					entries.push({
						"type" : "app",
						data : app
					});
				}
			}
		}

		// clear all the list and grid entries in the HTML
		$(".subcontent_grid").html("");
		$(".subcontent_list").html("");

		var theme = device.theme_id;

		var html_grid = $(".subcontent_grid");
		var html_list = $(".subcontent_list");

		var grid_class_suffix = "_" + device.grid_columns;

		var current_row;

		for ( var i = 0; i < entries.length; i++) {

			// for the grid, put in cells
			var entry = entries[i];

			var grid_image1 = "";
			var grid_image2 = "";

			var list_image1 = "";
			var list_image2 = "";
			var text = "";

			var identifier = "";
			var fn = null;

			var id_name = "";
			var id_value = "";

			if (entry.type == "app") {

				id_name = "app_id";
				id_value = entry.data;

				grid_image1 = App.Constants.THEME_URL + theme
						+ "/images/tile.png";
				grid_image2 = App.Constants.THEME_URL + theme + "/apps/"
						+ entry.data + ".png";

				list_image1 = App.Constants.THEME_URL + theme
						+ "/images/tile_list.png";
				list_image2 = App.Constants.THEME_URL + theme + "/apps/"
						+ entry.data + ".png";

				text = entry.data;
				identifier = "app_id=\"" + entry.data + "\"";
				fn = App.View.onAppButtonClicked;
			} else if (entry.type == "contact") {
				var contact_id = entry.data;
				var contact = App.Model.getContactById(contact_id);

				id_name = "contact_id";
				id_value = contact_id;

				if (contact != null) {

					grid_image1 = App.Constants.THEME_URL + theme
							+ "/images/tile.png";
					grid_image2 = App.Constants.THEME_URL + theme
							+ "/contacts/" + contact.contact_class + ".png";

					list_image1 = App.Constants.THEME_URL + theme
							+ "/images/tile_list.png";
					list_image2 = App.Constants.THEME_URL + theme
							+ "/contacts/" + entry.data + ".png";
					text = contact.contact_name;
					identifier = "contact_id=\"" + contact.contact_id + "\"";
					fn = App.View.onContactButtonClicked;
				} else {
					continue;
				}
			}

			if (i % device.grid_columns == 0 || current_row == null) {
				// start row
				current_row = $("<tr>");
				html_grid.append(current_row);
			}

			// create cell
			var grid_class = device.theme_id + " grid_entry grid_person person";
			var grid_entry = $("<td class='phone_cell'><span " + identifier
					+ " class='" + grid_class + "'>"
					+ "<img class=\"person-image image-up\" src=\""
					+ grid_image1 + "\"/>" + "<span class=\"person-text\">"
					+ text + "</span>" + "</span></td>");

			var template_data = {
				"attribute_name" : id_name,
				"attribute_value" : id_value,
				"background" : grid_image1,
				"icon" : grid_image2,
				"text" : text,
				"suffix_class" : grid_class_suffix
			};

			var grid_entry_template = _.template($("#template_grid_entry")
					.html());
			var grid_entry = grid_entry_template(template_data);
			grid_entry = $(grid_entry);

			current_row.append(grid_entry);
			grid_entry.click(fn);

			// for the list, just add it
			var list_class = device.theme_id + " list_entry list_person person";

			var list_entry = $("<li><span " + identifier + " class='"
					+ list_class + "'>"
					+ "<img class=\"list_person_image image-up\" src=\""
					+ list_image1 + "\"/>"
					+ "<span class=\"list_person_text\">" + text + "</span>"
					+ "</span></li>");

			var template_data = {
				"attribute_name" : id_name,
				"attribute_value" : id_value,
				"background" : list_image1,
				"icon" : list_image2,
				"text" : text
			};
			var list_entry_template = _.template($("#template_list_entry")
					.html());
			var list_entry = list_entry_template(template_data);
			list_entry = $(list_entry);

			list_entry.click(fn);

			html_list.append(list_entry);

		}
		
		

		App.View.redrawApps();
		App.View.redrawContacts();

		App.View.setCurrentTheme(device.theme_id);

	},

	showDragDrop : function() {
		Log.d("App.View.showDragDrop");

		$("#content_phone").show();
		$(".subcontent_grid").hide();
		$(".subcontent_list").hide();
		$("#grid_columns_view").hide();
		$(".subcontent_dragdrop").show();
		$(".subcontent_dragdrop_right").show();

		// $(".subcontent_dragdrop").html("");
		$(".dragdrop_table").html("");
		var deviceLayout = App.Model.getCurrentDeviceLayout();
		var max_rows = 10;
		var page_id = deviceLayout.getCurrentPageId();

		for ( var row_index = 1; row_index < max_rows; row_index++) {

			var left_obj = {
				row : row_index,
				position : "LEFT",
				"page_id": page_id
			};
			var right_obj = {
				row : row_index,
				position : "RIGHT",
				"page_id": page_id
			};
			var banner_obj = {
				row : row_index,
				position : "BANNER",
				"page_id": page_id
			};

			var left_cell = deviceLayout.getCellAt(left_obj);
			var right_cell = deviceLayout.getCellAt(right_obj);
			var banner_cell = deviceLayout.getCellAt(banner_obj);

			var empty_left_cell = $("<td class='dnd_cell dnd_left_cell dnd_empty_cell'></td>");
			var empty_right_cell = $("<td class='dnd_cell dnd_right_cell dnd_empty_cell'></td>");

			if (left_cell || right_cell || banner_cell) {
				var row = $("<tr></tr>");
				$(".dragdrop_table").append(row);
				if (banner_cell) {
					banner_cell.render(row);
					// var content = "<tr><td cell_id='" + banner_cell.getId() +
					// "' class='dnd_cell dnd_banner_cell' colspan='2'>banner "
					// + banner_cell.type + "</td></tr>";
					// $(".dragdrop_table").append(content);

				} else if (left_cell && right_cell) {
					left_cell.render(row);
					right_cell.render(row);

				} else if (left_cell) {
					left_cell.render(row);
					row.append(empty_right_cell);
					// var content = "<tr><td cell_id='" + left_cell.getId() +
					// "' class='dnd_cell dnd_left_cell' >left " +
					// left_cell.type + "</td><td class='dnd_cell
					// dnd_right_cell'/></tr>";
					// $(".dragdrop_table").append(content);

				} else if (right_cell) {
					row.append(empty_left_cell);
					right_cell.render(row);
					// var content = "<tr><td class='dnd_cell
					// dnd_left_cell'/><td cell_id='" + right_cell.getId() + "'
					// class='dnd_cell dnd_right_cell'>right " + right_cell.type
					// + "</td></tr>";
					// $(".dragdrop_table").append(content);

				}

			} else {
				break;
			}

		}

		$(".dnd_cell").unbind();
		$(".dnd_cell").click(function() {
			var cell = $(this);
			var cell_id = cell.attr("cell_id");
			if (cell_id) {
				var deviceLayout = App.Model.getCurrentDeviceLayout();
				var page_id = deviceLayout.getCurrentPageId();
				var cell_e = deviceLayout.getCellById(cell_id, page_id);
				var active_cell = deviceLayout.getSelectedCell();
				if (active_cell == cell_e) {
					active_cell.setSelected(false);
				} else if (active_cell != null) {
					active_cell.setSelected(false);
					cell_e.setSelected(true);
				} else {
					cell_e.setSelected(true);
				}

				App.View.showDragDrop();

			}
		});
		
		$(".page_cell").unbind();
		$(".page_cell").click(function() {
			var cell = $(this);
			var page_id = cell.attr("page_id");
			if ( page_id == "add" || page_id == "blank" ) {
				return;
			} else {
				var deviceLayout = App.Model.getCurrentDeviceLayout();
				deviceLayout.setCurrentPageId(page_id);
				App.View.showDragDrop();
			}
		});

		$("#dragdrop_arrow_left").unbind();
		$("#dragdrop_arrow_left").click(function() {
			var layout = App.Model.getCurrentDeviceLayout();
			layout.moveSelectedCellLeft();
		});
		$("#dragdrop_arrow_right").unbind();
		$("#dragdrop_arrow_right").click(function() {
			var layout = App.Model.getCurrentDeviceLayout();
			layout.moveSelectedCellRight();
		});
		$("#dragdrop_arrow_up").unbind();
		$("#dragdrop_arrow_up").click(function() {
			var layout = App.Model.getCurrentDeviceLayout();
			layout.moveSelectedCellUp();
		});
		$("#dragdrop_arrow_down").unbind();
		$("#dragdrop_arrow_down").click(function() {
			var layout = App.Model.getCurrentDeviceLayout();
			layout.moveSelectedCellDown();
		});
		$("#dragdrop_arrow_remove").unbind();
		$("#dragdrop_arrow_remove").click(function() {
			var layout = App.Model.getCurrentDeviceLayout();
			layout.removeSelectedCell();
		});

		$(".button_setting").unbind();
		$(".button_setting").click(function() {
			var clicked = $(this);
			var name = clicked.attr("setting_key");
			var value = clicked.attr("value");
			var device_id = App.Model.getCurrentDevice().device_id;
			App.Model.getCurrentDeviceLayout().updateSetting(name, value);
			
			if ( name == "customisable" ) {
				if ( value == "yes" ) {
					$(".column_left").show();
				} else {
					$(".column_left").hide();
					var deviceLayout = App.Model.getCurrentDeviceLayout();
					deviceLayout.setCurrentPageId(page_id);
					App.View.showDragDrop();
					
				}
			}
			// App.Server.updateDeviceSetting(device_id, name, value);
		});

		// TOTO FIND SYNTAX FOR UPDATING NAMED ATTRIBUES
		// AND DO IT
		// $(".button_setting").each(function(btn) {

		var keys = [ "customisable", "send_gps", "dragdrop", "haptic", "show_version" ];
		for ( var index = 0; index < keys.length; index++) {
			var key = keys[index];
			var currentValue = App.Model.getCurrentDeviceLayout().getSettingValueOrDefault(key, "no");
			var entries = $(".button_setting[setting_key='" + key + "']");
			var btn1 = $(entries[0]);
			var btn2 = $(entries[1]);
			var yesKey = btn1.attr("value") == "yes" ? btn1 : btn2;
			var noKey = btn1.attr("value") == "yes" ? btn2 : btn1;
			if (currentValue == "yes") {
				yesKey.addClass("active");
				noKey.removeClass("active");
			} else {
				yesKey.removeClass("active");
				noKey.addClass("active");
			}
		}
		
		var customisable = App.Model.getCurrentDeviceLayout().getSettingValueOrDefault("customisable", "no");
		if ( customisable == "yes" ) {
			$(".column_left").show();
		} else {
			$(".column_left").hide();
		}
		var currentPageId = App.Model.getCurrentDeviceLayout().getCurrentPageId();
		$(".page_cell").removeClass("page_cell_selected");
		$(".page_cell[page_id='" + currentPageId + "']").addClass("page_cell_selected");
		
		// $(".button_setting[setting_key='foo']");
		// attr("value") yes|no
		// $(".button_setting[setting_key='foo']");

		$(".button_list").attr("src", "/img/img_v2/btn_list_off.png");
		$(".button_grid").attr("src", "/img/img_v2/btn_grid_off.png");

		$(".button_grid_size1")
				.attr("src", "/img/img_v2/btn_grid_1col_off.png");
		$(".button_grid_size2")
				.attr("src", "/img/img_v2/btn_grid_2col_off.png");
		$(".button_grid_size3")
				.attr("src", "/img/img_v2/btn_grid_3col_off.png");
		$(".button_grid_size4")
				.attr("src", "/img/img_v2/btn_grid_4col_off.png");

	},

	showGrid : function() {
		Log.d("App.View.showGrid");

		$("#content_phone").show();
		$(".subcontent_dragdrop").hide();
		$(".subcontent_dragdrop_right").hide();
		$(".subcontent_grid").show();
		$(".subcontent_list").hide();
		$("#grid_columns_view").show();

		var btn = $("#button_grid_list");
		// btn.attr("src", "/img/grid_list_list.png");
		btn.attr("src", "/img/img_v2/grid_list_grid.png");

		$(".button_list").attr("src", "/img/img_v2/btn_list_off.png");
		$(".button_grid").attr("src", "/img/img_v2/btn_grid_on.png");

		var device = App.Model.getCurrentDevice();
		if (device != null) {
			var cols = device.grid_columns;
			if (cols == null) {
				cols = 3;
			}

			$(".button_grid_size2").attr("src",
					"/img/img_v2/btn_grid_2col_off.png");
			$(".button_grid_size3").attr("src",
					"/img/img_v2/btn_grid_3col_off.png");
			$(".button_grid_size4").attr("src",
					"/img/img_v2/btn_grid_4col_off.png");

			var key = ".button_grid_size" + cols;
			var value = "/img/img_v2/btn_grid_" + cols + "col_on.png";
			$(key).attr("src", value);
		}

	},

	showList : function() {
		Log.d("App.View.showList");

		$("#content_phone").show();
		$(".subcontent_grid").hide();
		$(".subcontent_dragdrop").hide();
		$(".subcontent_dragdrop_right").hide();
		$(".subcontent_list").show();
		$("#grid_columns_view").hide();

		var btn = $("#button_grid_list");
		btn.attr("src", "/img/img_v2/grid_list_list.png");

		$(".button_list").attr("src", "/img/img_v2/btn_list_on.png");
		$(".button_grid").attr("src", "/img/img_v2/btn_grid_off.png");

		$(".button_grid_size1")
				.attr("src", "/img/img_v2/btn_grid_1col_off.png");
		$(".button_grid_size2")
				.attr("src", "/img/img_v2/btn_grid_2col_off.png");
		$(".button_grid_size3")
				.attr("src", "/img/img_v2/btn_grid_3col_off.png");
		$(".button_grid_size4")
				.attr("src", "/img/img_v2/btn_grid_4col_off.png");

	},

	hideGrid : function() {
		Log.d("App.View.hideGrid");

		$(".subcontent_grid").hide();
		$("#grid_columns_view").hide();
	},

	hideDragDrop : function() {
		Log.d("App.View.hideDragDrop");
		$(".subcontent_dragdrop").hide();
		$(".subcontent_dragdrop_right").hide();
	},

	hideList : function() {
		Log.d("App.View.hideList");
		$(".subcontent_list").hide();
	},

	onThemeButtonClicked : function() {
		Log.d("App.View.changeTheme");
		var t = $(this);
		var theme_id = t.attr("theme_id");
		App.Model.getCurrentDevice().theme_id = theme_id;
		App.Model.getCurrentDevice().theme_version = App.Constants.THEME_VERSION;

		// var device_id = App.Model.getCurrentDevice().device_id;
		App.Model.getCurrentDeviceLayout().updateSetting("theme_id", theme_id);
		App.Model.getCurrentDeviceLayout().updateSetting("theme_version", App.Constants.THEME_VERSION);

		// App.Model.getCurrentDeviceLayout().settings.theme_id = theme_id;
		// App.Model.getCurrentDeviceLayout().settings.theme_version =
		// App.Constants.THEME_VERSION;

		App.View.setCurrentTheme(theme_id);
		App.View.redrawDevice();
		App.View.redrawApps();
		App.View.redrawContacts();
		App.Model.getCurrentDevice().save();
		// App.Model.getCurrentDeviceLayout().save();
	},

	onAppButtonClicked : function() {
		Log.d("App.View.onAppButtonClicked");
		var t = $(this);
		var app_id = t.attr("app_id");

		App.Model.getCurrentDevice().toggleApp(app_id);
		
		var page_id = App.Model.getCurrentDeviceLayout().getCurrentPageId();
		if (!App.Model.getCurrentDeviceLayout().hasApp(app_id, page_id)) {
			App.Model.getCurrentDeviceLayout().addApp(app_id, page_id);
		}
//		App.Model.getCurrentDeviceLayout().toggleApp(app_id);
		App.Model.getCurrentDevice().save();
		App.View.redrawDevice();
		return;

		if (App.Model.getCurrentDevice().hasApp(app_id)) {

			// prompt to remove it
			var options = {
				header : "Remove App?",
				body : "This will remove '" + app_id
						+ "' from this device.  Do you really want to?",
				yesButton : true,
				noButton : true,
				onYes : function() {
					App.Model.getCurrentDevice().toggleApp(app_id);
					App.Model.getCurrentDeviceLayout().toggleApp(app_id);
					App.Model.getCurrentDevice().save();
					App.View.redrawDevice();
				},
				onNo : function() {

				}
			};
			Common.showModal(options);
		} else {
			App.Model.getCurrentDevice().toggleApp(app_id);
			App.Model.getCurrentDeviceLayout().toggleApp(app_id);
			App.Model.getCurrentDevice().save();
			App.View.redrawDevice();
		}
	},

	onContactButtonClicked : function() {
		Log.d("App.View.onContactButtonClicked");
		var t = $(this);
		var contact_id = t.attr("contact_id");

		App.Model.getCurrentDevice().toggleContact(contact_id);
		
		var page_id = App.Model.getCurrentDeviceLayout().getCurrentPageId();
		if (!App.Model.getCurrentDeviceLayout().hasContact(contact_id, page_id)) {
			App.Model.getCurrentDeviceLayout().addContact(contact_id, page_id);
		}
		
//		App.Model.getCurrentDeviceLayout().toggleContact(contact_id);
		App.View.redrawDevice();
		App.Model.getCurrentDevice().save();
		return;

		if (App.Model.getCurrentDevice().hasContact(contact_id)) {
			// prompt to remove it
			var options = {
				header : "Remove Contact?",
				body : "This will remove '" + contact_id
						+ "' from this device.  Do you really want to?",
				yesButton : true,
				noButton : true,
				onYes : function() {
					App.Model.getCurrentDevice().toggleContact(contact_id);
					App.Model.getCurrentDeviceLayout()
							.toggleContact(contact_id);
					App.View.redrawDevice();
					App.Model.getCurrentDevice().save();
				},
				onNo : function() {

				}
			};
			Common.showModal(options);
		} else {
			App.Model.getCurrentDevice().toggleContact(contact_id);
			App.Model.getCurrentDeviceLayout().toggleContact(contact_id);
			App.View.redrawDevice();
			App.Model.getCurrentDevice().save();
		}

	},

	/**
	 * displays the 'add a new contact' ui
	 */
	drawAddContactPage : function() {
		Log.d("App.View.drawAddContactPage");

		var node = $("#dialog");
		node.html("");

		var keyPairs = {};
		var html = _.template($("#template_addcontact").html());
		var modalNode = $(html(keyPairs));
		node.append(modalNode);
		modalNode.modal();
		$("#contact_name").focus();

		var loginButton = $("#button_login");
		$("#button_add_contact_cancel").unbind().click(function() {
			$("#dialog").html();
		});
		$("#button_add_contact")
				.unbind()
				.click(
						function() {
							var contact_name = $("#contact_name").val();
							var contact_number = $("#contact_name").val();
							var contact_email = $("#contact_name").val();
							var contact_class = $("#contact_class").val();
							Server
									.contact_add(
											contact_name,
											contact_number,
											contact_email,
											contact_class,
											{
												onSuccess : function() {
													App.Server.fetchAndRedrawContacts();
													$("#add_contact_modal").remove();
												},
												onFailure : function() {
													Log.d("drawAddContactPage: failure on contact_add");
													$("#add_contact_modal").remove();
												}
											});
							$("#dialog").html();
						});
		return;

	},

	onPhoneButtonClicked : function() {

		Log.d("App.View.onPhoneButtonClicked");
		var t = $(this);
		var device_id = t.attr("device_id");

		// fetches teh configuraton from the server
		// onSuccess it will redraw the screen
		App.Server.getDeviceConfig(device_id);

	},

	onMapButtonClicked : function() {
		Log.d("App.View.onMapButtonClicked");
		if (App.Model.getCurrentDevice() == null) {
			return;
		}

		var btn = $("#button_map");
		if (App.Model.getDisplayMode() == "phone") {
			// flip to map
			App.Model.setDisplayMode("map");
			btn.attr("src", "/img/location_pressed.png");
			$(".subcontent_map").show();
			App.View.hideGrid();
			App.View.hideList();

		} else {
			// flip to phone
			App.Model.setDisplayMode("phone");
			btn.attr("src", "/img/location_unpressed.png");
			$(".subcontent_map").hide();
			if (App.Model.getCurrentDevice().displayMode == "grid") {
				App.View.showGrid();
			} else {
				App.View.showList();
			}
		}

	},

	/**
	 * flips phone between grid and list
	 */
	onToggleGridListButtonClicked : function() {
		Log.d("App.View.onToggleGridListButtonClicked");
		if (App.Model.getCurrentDevice() == null) {
			return;
		}

		var mapBtn = $("#button_map");
		mapBtn.attr("src", "/img/location_unpressed.png");

		$(".subcontent_map").hide();
		if (App.Model.getCurrentDevice().displayMode == "grid") {
			// flip to list
			App.Model.getCurrentDevice().displayMode = "list";
			App.Model.getCurrentDevice().save();
			App.View.showList();

		} else if (App.Model.getCurrentDevice().displayMode == "draddrop") {
			// flip to dragdrop
			App.Model.getCurrentDevice().displayMode = "dragdrop";
			App.Model.getCurrentDevice().save();
			App.View.showDragDrop();

		} else {
			// flip to grid
			App.Model.getCurrentDevice().displayMode = "grid";
			App.Model.getCurrentDevice().save();
			App.View.showGrid();

		}
	},

	onDragDropButtonClicked : function() {
		App.View.setPhoneAsDragDropView();
	},
	onGridButtonClicked : function() {
		App.View.setPhoneAsGridView();
	},
	onListButtonClicked : function() {
		App.View.setPhoneAsListView();
	},

	onGridSizeButtonClicked : function() {

		var node = $(this);
		var device = App.Model.getCurrentDevice();

		var device_id = device.device_id;
		var val = node.attr("cols");
		// $("#grid_columns").val();
		if (val != null) {
			device.grid_columns = val;
			// $(".grid_columns_text").html("Columns (" + val + ")");
			App.Server.saveGridColumns(device_id, val);
			App.View.setPhoneAsGridView();
			App.View.redrawDevice();
		}
	},

	setPhoneAsDragDropView : function() {
		Log.d("App.View.setPhoneAsDragDropView");
		if (App.Model.getCurrentDevice() == null) {
			return;
		}

		if (App.Model.getCurrentDevice().displayMode == "dragdrop")
			return;
		var mapBtn = $("#button_map");
		mapBtn.attr("src", "//img/location_unpressed.png");
		// flip to grid
		App.Model.getCurrentDevice().displayMode = "dragdrop";
		App.Model.getCurrentDevice().save();
		App.View.showDragDrop();
	},

	setPhoneAsGridView : function() {
		Log.d("App.View.setPhoneAsGridView");
		if (App.Model.getCurrentDevice() == null) {
			return;
		}

		if (App.Model.getCurrentDevice().displayMode == "grid")
			return;
		var mapBtn = $("#button_map");
		mapBtn.attr("src", "//img/location_unpressed.png");
		// flip to grid
		App.Model.getCurrentDevice().displayMode = "grid";
		App.View.showGrid();
		App.Model.getCurrentDevice().save();
	},

	setPhoneAsListView : function() {
		Log.d("App.View.setPhoneAsListView");
		if (App.Model.getCurrentDevice() == null) {
			return;
		}

		if (App.Model.getCurrentDevice().displayMode == "list")
			return;
		var mapBtn = $("#button_map");
		mapBtn.attr("src", "/img/location_unpressed.png");

		$(".subcontent_map").hide();
		// flip to list
		App.Model.getCurrentDevice().displayMode = "list";
		App.View.showList();
		App.Model.getCurrentDevice().save();
	},


	onKeyPress: function(keyEvent) {

		var charCode = keyEvent.charCode;
		var keyCode = keyEvent.charCode;
		var altKey = keyEvent.altKey;
		var ctrlKey = keyEvent.ctrlKey;
		var shiftKey = keyEvent.shiftKey;
		var isLeftArrow = charCode == 0 && keyCode == 37;
		var isRightArrow = charCode == 0 && keyCode == 39;
		var isUpArrow = charCode == 0 && keyCode == 38;
		var isDownArrow = charCode == 0 && keyCode == 40;
		var isEnter = charCode == 0 && keyCode == 13;
		var isTab = charCode == 0 && keyCode == 9;

		var is1 = charCode == 49 && keyCode == 0;
		var is2 = charCode == 50 && keyCode == 0;
		var is3 = charCode == 51 && keyCode == 0;
		var is4 = charCode == 52 && keyCode == 0;
		var is5 = charCode == 53 && keyCode == 0;
		var is6 = charCode == 54 && keyCode == 0;
		var is7 = charCode == 55 && keyCode == 0;
		var is8 = charCode == 56 && keyCode == 0;
		var is9 = charCode == 57 && keyCode == 0;

	}

};