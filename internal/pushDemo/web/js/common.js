var urlInput = document.getElementById("url"),
    groupInput = document.getElementById("group"),
    userInput = document.getElementById("user"),
    bsidInput = document.getElementById("bsid");

init();

function init() {
    var url = sessionStorage[prefix + "url"] || "",
        group = sessionStorage[prefix + "group"] || "",
        user = sessionStorage[prefix + "user"] || "",
        bsid = sessionStorage[prefix + "bsid"] || "";

    urlInput.value = url;
    groupInput.value = group;
    userInput.value = user;
    bsidInput.value = bsid;
}

function setToLocalStorage(url, group, user, bsid) {
    sessionStorage[prefix + "url"] = url;
    sessionStorage[prefix + "group"] = group;
    sessionStorage[prefix + "user"] = user;
    sessionStorage[prefix + "bsid"] = bsid;
}

function getFinalUrl() {
    var rst = "",
        err = "";

    var url = urlInput.value,
        group = groupInput.value,
        user = userInput.value,
        bsid = bsidInput.value;

    if (!url) {
        err = "Please enter url!!!";
        return {rst, err};
    }
    var index = url.indexOf("?");
    if (index !== -1) {
        url = url.substring(0, index);
    }

    var queryString;
    if (!group && !user && !bsid) {
        queryString = "";
    } else {
        queryString = `?group=${encodeURIComponent(group)}&user=${encodeURIComponent(user)}&bsid=${encodeURIComponent(bsid)}`
    }

    setToLocalStorage(url, group, user, bsid);

    rst = url + queryString
    return {rst, err};
}