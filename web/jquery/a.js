window.onload = function () {
    console.info("===")

    var jsonParams = {
            method: 1,
            params: {
                name: "test测试",
                age: 20
            }
        },
        jsonStr = JSON.stringify(jsonParams);

    $.ajax({
        url: "http://127.0.0.1/test",
        type: "POST",
        // contentType: 'application/json',
        // contentType: 'application/x-www-form-urlencoded',
        data: {
            jsonParams: encodeURIComponent(jsonStr)
        },
        dataType: "text",
        async: false,
        general: false,
        success: function (data) {
            console.info(data);
        },
        error: function (data) {
            console.error(data);
        }
    });
};

