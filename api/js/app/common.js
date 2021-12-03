/**
 * general purpose logger
 * 
 * @class
 */
var Log = {

    DEBUG_ENABLED: true,
    ERROR_ENABLED: true,
    WARN_ENABLED: true,
    INFO_ENABLED: true,

	/**
	 * logs some debug information
	 */
	d : function(msg) {
        if ( this.DEBUG_ENABLED ) {
    		console.log(msg);
        }
	},

  	/**
	 * logs some debug information
	 */
	i: function(msg) {
        if ( this.INFO_ENABLED ) {
    		console.log(msg);
        }
	},

	/**
	 * logs an exception
	 */
	e: function(error) {
        if ( this.ERROR_ENABLED ) {
    		console.error(error);
        }
	},

	w: function(msg) {
        if ( this.WARN_ENABLED ) {
    		console.warn(msg);
        }
	},

};

/**
 * utility methods
 */
var Common = {

	init: function() {
	},
	
	/**
	 * simple defer function
	 */
	run : function(fn) {
		setTimeout(fn, 1);
	},
	
	/**
	 * executes a callback in the passed number of millis
	 */
	delay : function(callback, millis) {
		setTimeout(callback, millis);
	},	

	is_iPad : function() {
		return navigator.userAgent.match(/iPad/i) != null;
	},

	is_iPhone : function() {
		return (
		// Detect iPhone
		(navigator.platform.indexOf("iPhone") != -1) ||
		// Detect iPod
		(navigator.platform.indexOf("iPod") != -1));
	},

	Clipboard : {
		store : function(key, value) {
			if (window.clipboardData) {
				window.clipboardData.setData(key, value);
			}
		}
	},

	/**
	 * the Poller is used to kick off named scripts so I don't have to remember
	 * pollerIds
	 * 
	 * Usage:
	 *  // to start Common.Poller.start("fred", 10000, function() { Log.d("I
	 * have been called"); })
	 * 
	 *  // later on, to stop Common.Poller.stop("fred");
	 * 
	 */
	Poller : {

		// holds the pollerids mapped against the provided id
		state : {},

		/**
		 * started a named poller, every "interval" millis, calling the callback
		 * if the id already exists, it ignores this request
		 */
		start : function(id, interval, callback) {
			if (!Common.Poller.state[id]) {
				var interval_id = setInterval(callback, interval);
				Common.Poller.state[id] = interval_id;
			}
		},

		stop : function(id) {
			var interval_id = Common.Poller.state[id];
			if (interval_id) {
				clearInterval(interval_id);
				delete Common.Poller.state[id];
			}
		},

	},

	Events : {
		subscribe : function(subscriptions) {
			for ( var index = 0; index < subscriptions.length; index++) {
				var subscription = subscriptions[index];
				var eventName = subscription[0];
				var callback = subscription[1];
				amplify.subscribe(eventName, callback);
			}
		},
		unsubscribe : function(subscriptions) {
			for ( var index = 0; index < subscriptions.length; index++) {
				var subscription = subscriptions[index];
				var eventName = subscription[0];
				var callback = subscription[1];
				amplify.unsubscribe(eventName, callback);
			}
		}

	},

	/**
	 * high level wrapper for persistent storage
	 */
	Store : {
		put : function(key, value) {
			amplify.store(key, value);
		},
		get : function(key, defaultValue) {
			var value = amplify.store(key);
			return value || defaultValue;
		},
		del : function(key) {
			amplify.store(key, null);
		},
		contains : function(key) {
			return amplify.store(key);
		}
	},

	/**
	 * collection of useful cross-app GUI utilities
	 */
	GUI : {

		notifyX: function(title, text) {
			
			Messenger.options = {
					extraClasses: 'messenger-fixed messenger-on-top messenger-on-right',
					theme: 'ice'
			};
		
			Messenger().post(title + text);
		},
		
		// $.pnotify.defaults.styling = "bootstrap";
		// see _init for the basic setup
		notify : function(title, text) {
			Log.d("Common.GUI.notifyInfo()");
			$.pnotify({
				"title" : title,
				"text" : text
			});
		},

		notifyInfo : function(title, text) {
			Log.d("Common.GUI.notifyInfo()");
			$.pnotify({
				"title" : title,
				"text" : text,
				type : "info"
			});
		},

		notifyInfoSticky : function(title, text) {
			Log.d("Common.GUI.notifyInfo()");
			$.pnotify({
				"title" : title,
				"text" : text,
				type : "info",
				hide : false,
				sticker : false
			});
		},

		notifyError : function(title, text) {
			Log.d("Common.GUI.notiftyError()");
			$.pnotify({
				"title" : title,
				"text" : text,
				type : "error",
				hide : true
			});
		},

		notifyErrorSticky : function(title, text) {
			Log.d("Common.GUI.notiftyErrorSticky()");
			$.pnotify({
				"title" : title,
				"text" : text,
				type : "error",
				hide : false,
				sticker : false
			});
		},

		notifySuccess : function(title, text) {
			Log.d("Common.GUI.notifySuccess()");
			$.pnotify({
				"title" : title,
				"text" : text,
				type : "success"
			});
		},

		notifySuccessSticky : function(title, text) {
			Log.d("Common.GUI.notifySuccess()");
			$.pnotify({
				"title" : title,
				"text" : text,
				type : "success",
				hide : false,
				sticker : false
			});
		},

		prompt : function(data) {
			// Common.GUI.prompt_avgrund(data);
			Common.GUI.prompt_bootstrap(data);
		},

		onKeyPressEnter : function(control, callback) {
			control.keypress(function(event) {
				if (event.which == 13) {
					callback();
				}
			});
		},

		onKeyPressTab : function(control, callback) {
			control.keypress(function(event) {
				if (event.which == 0) {
					callback();
					return false;
				}
			});
		},

		/**
		 * displays a popup prompt and acts on the data passed in
		 */
		prompt_bootstrap : function(data) {
			var modalNode = $("#modal");
			modalNode.find("#modal-title").html(data.title);
			modalNode.find(".modal_description").html(data.description);
			if (!data.yesButton && !data.noButton) {
				modalNode.find(".button_modal_yes").html("Ok");
				modalNode.find(".button_modal_yes").show();
				modalNode.find(".button_modal_no").hide();

			} else {
				modalNode.find(".button_modal_yes").show();
				modalNode.find(".button_modal_no").show();

				modalNode.find(".button_modal_yes").html(data.yesButton);
				modalNode.find(".button_modal_no").html(data.noButton);
				modalNode.find(".button_modal_yes").unbind().click(data.onYes);
				modalNode.find(".button_modal_no").unbind().click(data.onNo);
			}
			$("#modal").modal();
		},

		prompt_avgrund : function(data) {

			// var modalNode = $("#modal");
			$(".button_logout")
					.avgrund(
							{
								height : 200,
								holderClass : 'custom',
								showClose : true,
								showCloseText : 'Close',
								enableStackAnimation : true,
								onBlurContainer : '.overall_container',
								template : '<p>So implement your design and place content here! If you want to close modal, please hit "Esc", click somewhere on the screen or use special button.</p>'
										+ '<div>'
										+ '<a href="http://github.com/voronianski/jquery.avgrund.js" target="_blank" class="github">Avgrund on Github</a>'
										+ '<a href="http://twitter.com/voronianski" target="_blank" class="twitter">Twitter</a>'
										+ '<a href="http://dribbble.com/voronianski" target="_blank" class="dribble">Dribbble</a>'
										+ '</div>'

							});
		},

    },

    Strings: {

		/**
		 * Convert number of bytes into human readable format
		 * 
		 * @param integer
		 *            bytes Number of bytes to convert
		 * @param integer
		 *            precision Number of digits after the decimal separator
		 * @return string
		 */
		friendly_bytesToSize : function(bytes, precision) {
			var kilobyte = 1024;
			var megabyte = kilobyte * 1024;
			var gigabyte = megabyte * 1024;
			var terabyte = gigabyte * 1024;

			if ((bytes >= 0) && (bytes < kilobyte)) {
				return bytes + ' B';

			} else if ((bytes >= kilobyte) && (bytes < megabyte)) {
				return (bytes / kilobyte).toFixed(precision) + ' KB';

			} else if ((bytes >= megabyte) && (bytes < gigabyte)) {
				return (bytes / megabyte).toFixed(precision) + ' MB';

			} else if ((bytes >= gigabyte) && (bytes < terabyte)) {
				return (bytes / gigabyte).toFixed(precision) + ' GB';

			} else if (bytes >= terabyte) {
				return (bytes / terabyte).toFixed(precision) + ' TB';

			} else {
				return bytes + ' B';
			}

		},

		formatMillis : function(totalTime) {

			var remainingTime = totalTime;
			var SECOND = 1000;
			var MINUTE = SECOND * 60;
			var HOUR = MINUTE * 60;
			var DAY = HOUR * 24;
			var formatted = "";

			if (remainingTime >= DAY) {
				var remainder = remainingTime % DAY;
				var days = ((remainingTime - remainder) / DAY);
				remainingTime = remainder;
				formatted += (days + "d " );
			}

			if (remainingTime >= HOUR) {
				var remainder = remainingTime % HOUR;
				var hours = ((remainingTime - remainder) / HOUR).toFixed(0);
				remainingTime = remainder;
				formatted += (hours + "h " )
			}

			if (remainingTime >= MINUTE) {
				var remainder = remainingTime % MINUTE;
				var minutes = ((remainingTime - remainder) / MINUTE).toFixed(0);
				remainingTime = remainder;
				formatted += (minutes + "m " );
			}

			if (remainingTime >= SECOND) {
				var remainder = remainingTime % SECOND;
				var seconds = ((remainingTime - remainder) / SECOND).toFixed(0);
				remainingTime = remainder;
				formatted += (seconds + "s " );
			}

			if (formatted == "" && remainingTime > 0) {
				formatted += (remainingTime + "ms");
			}
			return formatted;

		},

		formatDayHourMinute: function(totalTimeInSeconds) {

			var remainingTime = totalTimeInSeconds;
			var SECOND = 1;
			var MINUTE = SECOND * 60;
			var HOUR = MINUTE * 60;
			var DAY = HOUR * 24;
			var formatted = "";

			if (remainingTime >= DAY) {
				var remainder = remainingTime % DAY;
				var days = ((remainingTime - remainder) / DAY);
				remainingTime = remainder;
				formatted += (days + "d " );
			}

			if (remainingTime >= HOUR) {
				var remainder = remainingTime % HOUR;
				var hours = ((remainingTime - remainder) / HOUR).toFixed(0);
				remainingTime = remainder;
				formatted += (hours + "h " )
			}

			if (remainingTime >= MINUTE) {
				var remainder = remainingTime % MINUTE;
				var minutes = ((remainingTime - remainder) / MINUTE).toFixed(0);
				remainingTime = remainder;
				formatted += (minutes + "m ");
			}

			return formatted;
		
		},

		formatSeconds: function(totalTime) {
			return Common.GUI.formatMillis(totalTime*1000);
		},

		friendlyNumber: function(num_in) {
			return $.number(num_in, 3, ',');
		},
		
		/**
		 * formats an epoch (seconds, not millis) to HH:MM:SS
		 */
		formatEpoch : function(time, formatType) {
			if (time == null || time == 0) {
				return "";
			}
			try {
				var d = new Date(time);
				var h = d.getHours();
				var m = d.getMinutes();
				var s = d.getSeconds();
				h = h < 10 ? "0" + h : h;
				m = m < 10 ? "0" + m : m;
				s = s < 10 ? "0" + s : s;

				return h + ":" + m + ":" + s;
				// return d.toString(formatType);
			} catch (err) {
				return "";
			}
		},

		/**
		 * loads and populates some html based on a template selector and some data
		 */
		template_html: function(template_class, template_data) {
			var ref_html = $(template_class).html();
			var template = _.template(ref_html);
			var html = template(template_data);
			return html;
		},
		
			
	},

	
};

String.prototype.endsWith = function(suffix) {
	return this.indexOf(suffix, this.length - suffix.length) !== -1;
};

