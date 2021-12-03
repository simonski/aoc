var App = App || {};
App.Objects = {

	User: {
		isAdmin: function() {
			return this.role == "admin";
		},
		getName: function() {
			return this.user;
		},
	},	

	Device: {

	},
	
};