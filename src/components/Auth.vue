<template>
  <div class="auth">
    <!-- <button @click.prevent="authflow()" class="button is-default">Start</button> -->
    <button class="button is-gray">Soon coming</button>
  </div>
</template>

<script>
import { mapState, mapActions, mapGetters, mapMutations } from "vuex";

const axios = require("axios");

const electron = require("electron");
const remote = electron.remote;
export default {
  name: "Auth",
  data: () => {
    return {
      CLIENT_ID: "f3a2c34ebdced1411fee"
    };
  },
  methods: {
    requestGithubToken(code) {},
    handleCallback(url) {},
    authflow() {
      console.log("authflow");
      // Your GitHub Applications Credentials
      var options = {
        client_id: this.CLIENT_ID,
        scopes: ["user:email", "notifications"] // Scopes limit access for OAuth tokens.
      };

      // Build the OAuth consent page URL
      var authWindow = new remote.BrowserWindow({
        width: 800,
        height: 600,
        show: false,
        "node-integration": false
      });
      var githubUrl = "https://github.com/login/oauth/authorize?";
      var authUrl =
        githubUrl +
        "client_id=" +
        options.client_id +
        "&scope=" +
        options.scopes;
      authWindow.loadURL(authUrl);
      authWindow.show();

      // Handle the response from GitHub - See Update from 4/12/2015
      authWindow.webContents.on("will-navigate", function(event, url) {
        console.log(url);
        var raw_code = /code=([^&]*)/.exec(url) || null;
        var code = raw_code && raw_code.length > 1 ? raw_code[1] : null;
        var error = /\?error=(.+)$/.exec(url);

        if (code || error) {
          // Close the browser if code found or error
          authWindow.destroy();
        }

        // If there is a code, proceed to get token from github
        if (code) {
          axios({
            url: "https://github.com/login/oauth/access_token",
            method: "post",
            data: {
              client_id: this.CLIENT_ID
            }
          })
            .then(response => {
              console.log(response);
              window.localStorage.setItem(
                "githubtoken",
                response.body.access_token
              );
            })
            .catch(error => {
              console.error(error);
            });
        } else if (error) {
          alert(
            "Oops! Something went wrong and we couldn't" +
              "log you in using Github. Please try again."
          );
        }
      });

      authWindow.webContents.on("did-get-redirect-request", function(
        event,
        oldUrl,
        newUrl
      ) {
        console.log("newurl", newUrl);
        var raw_code = /code=([^&]*)/.exec(newUrl) || null;
        var code = raw_code && raw_code.length > 1 ? raw_code[1] : null;
        var error = /\?error=(.+)$/.exec(newUrl);

        if (code || error) {
          // Close the browser if code found or error
          authWindow.destroy();
        }

        // If there is a code, proceed to get token from github
        if (code) {
          axios({
            url: "https://github.com/login/oauth/access_token",
            method: "post",
            data: {
              client_id: this.CLIENT_ID
            }
          })
            .then(response => {
              console.log(response);
              window.localStorage.setItem(
                "githubtoken",
                response.body.access_token
              );
            })
            .catch(error => {
              console.error(error);
            });
        } else if (error) {
          alert(
            "Oops! Something went wrong and we couldn't" +
              "log you in using Github. Please try again."
          );
        }
      });

      // Reset the authWindow on close
      authWindow.on(
        "close",
        function() {
          authWindow = null;
        },
        false
      );
    }
  }
};
</script>

