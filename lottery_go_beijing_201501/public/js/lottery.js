var ws = new WebSocket("ws://localhost:4000/ws");
var status = -1;

$(document).ready(function () {
    ws.onopen = function () {
        $('#status').text("服务器已连接")
        $('button[name=start_btn]').show();
    };
    ws.onclose = function () {
        $('#status').text("服务已断开连接");
        $('button[name=start_btn]').hide();
        $('#stop_btn').hide();
    };
    ws.onmessage = function (msg) {
        if (status == 1)
            $('#current_info').text(msg.data);
        else if (status == 0) {
            $('#name').text(msg.data);
            $('table tbody').append('<tr><td>' + msg.data + '</td><td>' + $('#info').text() + '</td></tr>')
        }
    };

    $('button[name=start_btn]').click(function () {
        $('#start_panel').hide();
        $('#done_panel').hide();
        $('#doing_panel').show();
        ws.send("START");
        status = 1;
    });

    $('#stop_btn').click(function () {
        ws.send("STOP");
        status = 0;
        $('#doing_panel').hide();
        $('#done_panel').show();

        var info = $('#current_info').text();
        $('#info').text(info);
        ws.send(info);  // Get name.
    });
});
