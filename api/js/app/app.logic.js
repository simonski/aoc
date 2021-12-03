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
App.Logic = {

	init : function() {
		Log.d("App.Logic.init()");
//		window.onerror = App.handleError;
	},

	/**
	 * general purpose error handler (currently alerts only)
	 */
	handleError : function(err) {
		Log.e("App.Logic.handleError: unhandled error: " + err);
		alert("HandleError " + err);
	},

};
