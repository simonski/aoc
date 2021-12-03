var Device = function(data) {
	
    _.extend(this, data);
	
	this.save = function() {
			Log.d("Device.save");
			App.Server.saveDeviceConfig();
	};
			
	this.toggleApp = function(packageName) {
			Log.d("App.Phone.toggleApp(" + packageName + ")");
			if (!this.apps ) {
				this.apps = [];
			}
			if ( this.apps.indexOf(packageName) > -1 ) {
				var index = this.apps.indexOf(packageName);
				this.apps.splice(index, 1);
			} else {
				this.apps.push(packageName);
			}
	};
	
	this.hasContact = function(contact_id) {
		return this.contacts.indexOf(contact_id) > -1;
	};

	this.hasApp = function(app_id) {
		return this.apps.indexOf(app_id) > -1;
	};

	this.toggleContact = function(contact_id) {
			Log.d("App.Phone.toggleContact("+ contact_id + ")");
			if ( !this.contacts ) {
				this.contacts = [];
			}
			if ( this.contacts.indexOf(contact_id) > -1 ) {
				var index = this.contacts.indexOf(contact_id);
				this.contacts.splice(index, 1);
			} else {
				this.contacts.push(contact_id);
			}
	};

	this.removeContact = function(contact_id) {
			Log.d("App.Phone.removeContact("+ contact_id+")");
			if ( !this.contacts ) {
				this.contacts = [];
			}
			if ( this.contacts.indexOf(contact_id) > -1 ) {
				var index = this.contacts.indexOf(contact_id);
				this.contacts.splice(index, 1);
			}
	};

	this.addContact = function(contact_id) {
		Log.d("App.Phone.addContact(" + contact_id +")");
		if ( !this.contacts ) {
			this.contacts = [];
		}
		if ( this.contacts.indexOf(contact_id) === -1 ) {
			this.contacts.push(contact_id);
		}
		
	};
		
	this.addApp = function(packageName) {
			Log.d("App.Phone.addApp(" + packageName + ")");
			if (!this.apps ) {
				this.apps = [];
			}
			if ( this.apps.indexOf(packageName) === -1 ) {
				var index = this.apps.indexOf(packageName);
				this.apps.splice(index, 1);
			}
		};

	this.removeApp = function(packageName) {
			Log.d("App.Phone.removeApp(" + packageName + ")");
			if (!this.apps ) {
				this.apps = [];
			}
			if ( this.apps.indexOf(packageName) > -1 ) {
				var index = this.apps.indexOf(packageName);
				this.apps.splice(index, 1);
			}
	};

};
