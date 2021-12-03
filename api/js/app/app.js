/**
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
App.Constants = {
		THEME_URL : "/app/themes/",
		THEME_VERSION : "app",
		POLL_PERIOD : 2000,
		IS_IPAD: navigator.userAgent.match(/iPad/i) != null,
		IS_BROWSER: navigator.userAgent.match(/iPad/i) == null,
};

App.init = function() {
		
		Log.d("App.init()");

        amplify.subscribe(Server.Topic.LOGIN_REQUIRED, function() {
            Log.d("App.init(); Server.Topic.LOGIN_REQUIRED was broadcast, redrawing to welcome/login page.");
            App.View.drawLoginPage();
        });
        amplify.subscribe(Server.Topic.LOGIN_FAILURE, function() {
            Log.d("App.init(); Server.Topic.LOGIN_FAILURE was broadcast, redrawing to welcome/login page.");
            App.View.drawLoginPage();
        });
        amplify.subscribe(Server.Topic.LOGIN_SUCCESS, function() {
            Log.d("App.init(); Server.Topic.LOGIN_SUCCESS was broadcast, redrawing the logged in page.");
            App.View.drawLoggedInPage();
        });
        amplify.subscribe(Server.Topic.LOGOUT_SUCCESS, function() {
            Log.d("App.init(); Server.Topic.LOGOUT_SUCCESS was broadcast, redrawing the welcome/login page.");
            App.View.drawLoginPage();
        });
        amplify.subscribe(Server.Topic.REGISTER_SUCCESS, function() {
            Log.d("App.init(); Server.Topic.REGISTER_SUCCESS was broadcast, drawing the register success page.");
            App.View.drawRegisterSuccessPage();
        });
        amplify.subscribe(Server.Topic.REGISTER_FAILURE, function() {
            Log.d("App.init(); Server.Topic.REGISTER_FAILURE was broadcast, drawing the register failed page in page.");
            App.View.drawRegisterFailurePage();
        });

        Server.profile(null, {
            onSuccess: function() {
                Log.d("App.init() (Server.profile()) SUCCESS, draw logged in page.");
                App.View.drawLoggedInPage();
            },
            onFailure: function() {
                Log.d("App.init() (Server.profile()) FAILURE");
            }
        });

    App.View = {
        drawLoginPage: function() {
            Log.d("App.View.drawLoginPage() - the welcome, please login page");
        },
        drawLoggedInPage: function() {
            Log.d("App.View.drawLoggedInPage()");
        },
        drawRegisterSuccessPage: function() {
            Log.d("App.View.drawRegisterSuccess()");
        },
        drawRegisterFailurePage: function() {
            Log.d("App.View.drawRegisterFailurePage()");
        },
    }

};



