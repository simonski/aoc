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
App.SocketIO = {

    is_connected: false,

	init: function() {
		Log.d("App.SocketIO.init()");
	},

    connect: function() {
        if ( this.is_connected ) {
            console.log("socket.io: is_connected = true, returning");
            return false;
        } 
        console.log("socket.io: is_connected = false, trying to connect.");

        // var socket = io.connect('http://localhost:8080/');
        var socket = io.connect();
        App.SocketIO.socket = socket;

        socket.on('connect', function(msg) {
            console.log("socket.io: connect");
            this.is_connected = true;
        }),

        socket.on('disconnect', function(msg) {
            console.log("socket.io: disconnect");
            this.is_connected = false;
        }),

        socket.on('close', function(msg) {
            console.log("socket.io: close");
            this.is_connected = false;
        }),

        socket.on('error', function(error_msg) {
            console.error("socketio.error: " , error_msg);
            this.is_connected = false;
            socket.disconnect();

        });
            
        socket.on('news', function (data) {
            console.log(data);
            socket.emit('my other event', { my: 'data' });
        });

    },

    disconnect: function() {
        io.disconnect();
        if ( !this.is_connected ) {
            console.log("socketio: disconnect(); not connected.");
            return;
        }
        this.socket.emit('disconnect');
        this.socket = null;
        this.is_connected = false;

    },

};
