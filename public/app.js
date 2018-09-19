new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
        username: null, // Our username
        joined: false // True if email and username have been filled in
    },

    created: function() {
    },

    methods: {
        send: function () {
            if (this.newMsg != '') {
                this.ws.send(
                    JSON.stringify({
                        message: $('<p>').html(this.newMsg).text() // Strip out html
                    }
                ));
                this.newMsg = ''; // Reset newMsg
            }
        },

        join: function () {
            var self = this;
            this.username = $('<p>').html(this.username).text();
            if (!this.username) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            this.joined = true
            this.ws = new WebSocket('ws://' + window.location.host + '/ws?username='+this.username);
            this.ws.addEventListener('message', function(e) {
                var msg = JSON.parse(e.data);
                self.chatContent +=  '<div class="chip">'
                    + msg.username
                    + '</div>'
                    + emojione.toImage(msg.message) + '<br/>'; // Parse emojis

                var element = document.getElementById('chat-messages');
                element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
            });
        },
    }
});