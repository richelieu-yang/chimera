window.onload = function () {
    var text = `<html xmlns:v="urn:schemas-microsoft-com:vml"
xmlns:o="urn:schemas-microsoft-com:office:office"
xmlns:x="urn:schemas-microsoft-com:office:excel"
xmlns="http://www.w3.org/TR/REC-html40">

<head>
<meta http-equiv=Content-Type content="text/html; charset=utf-8">
<meta name=ProgId content=Excel.Sheet>
<meta name=Generator content="Microsoft Excel 15">
<link id=Main-File rel=Main-File
href="file:///C:/Users/lenovo/AppData/Local/Temp/msohtmlclip1/01/clip.htm">
<link rel=File-List
href="file:///C:/Users/lenovo/AppData/Local/Temp/msohtmlclip1/01/clip_filelist.xml">
<style>
<!--table
\t{mso-displayed-decimal-separator:"\\.";
\tmso-displayed-thousand-separator:"\\,";}
@page
\t{margin:.75in .7in .75in .7in;
\tmso-header-margin:.3in;
\tmso-footer-margin:.3in;}
.font5
\t{color:windowtext;
\tfont-size:9.0pt;
\tfont-weight:400;
\tfont-style:normal;
\ttext-decoration:none;
\tfont-family:等线;
\tmso-generic-font-family:auto;
\tmso-font-charset:134;}
tr
\t{mso-height-source:auto;
\tmso-ruby-visibility:none;}
col
\t{mso-width-source:auto;
\tmso-ruby-visibility:none;}
br
\t{mso-data-placement:same-cell;}
td
\t{padding-top:1px;
\tpadding-right:1px;
\tpadding-left:1px;
\tmso-ignore:padding;
\tcolor:black;
\tfont-size:11.0pt;
\tfont-weight:400;
\tfont-style:normal;
\ttext-decoration:none;
\tfont-family:等线;
\tmso-generic-font-family:auto;
\tmso-font-charset:134;
\tmso-number-format:General;
\ttext-align:general;
\tvertical-align:bottom;
\tborder:none;
\tmso-background-source:auto;
\tmso-pattern:auto;
\tmso-protection:locked visible;
\twhite-space:nowrap;
\tmso-rotate:0;}
.xl65
\t{text-align:center;
\tvertical-align:middle;}
.xl66
\t{text-align:center;
\tvertical-align:middle;
\tborder:.5pt solid windowtext;}
ruby
\t{ruby-align:left;}
rt
\t{color:windowtext;
\tfont-size:9.0pt;
\tfont-weight:400;
\tfont-style:normal;
\ttext-decoration:none;
\tfont-family:等线;
\tmso-generic-font-family:auto;
\tmso-font-charset:134;
\tmso-char-type:none;
\tdisplay:none;}
-->
</style>
</head>

<body link="#0563C1" vlink="#954F72">

<table border=0 cellpadding=0 cellspacing=0 width=868 style='border-collapse:
 collapse;width:651pt'>
<!--StartFragment-->
 <col width=124 span=7 style='mso-width-source:userset;mso-width-alt:3968;
 width:93pt'>
 <tr height=44 style='mso-height-source:userset;height:33.0pt'>
  <td height=44 class=xl66 width=124 style='height:33.0pt;width:93pt'>1</td>
  <td class=xl66 width=124 style='border-left:none;width:93pt'>2</td>
  <td class=xl66 width=124 style='border-left:none;width:93pt'>3</td>
  <td class=xl66 width=124 style='border-left:none;width:93pt'>4</td>
  <td class=xl66 width=124 style='border-left:none;width:93pt'>5</td>
  <td class=xl66 width=124 style='border-left:none;width:93pt'>6</td>
  <td class=xl66 width=124 style='border-left:none;width:93pt'>7</td>
 </tr>
 <tr height=44 style='mso-height-source:userset;height:33.0pt'>
  <td height=44 class=xl66 style='height:33.0pt;border-top:none'>测试1</td>
  <td class=xl66 style='border-top:none;border-left:none'>测试2</td>
  <td class=xl66 style='border-top:none;border-left:none'>测试3</td>
  <td class=xl66 style='border-top:none;border-left:none'>测试4</td>
  <td class=xl66 style='border-top:none;border-left:none'>测试5</td>
  <td class=xl66 style='border-top:none;border-left:none'>测试6</td>
  <td class=xl66 style='border-top:none;border-left:none'>测试7</td>
 </tr>
 <tr height=44 style='mso-height-source:userset;height:33.0pt'>
  <td height=44 class=xl66 style='height:33.0pt;border-top:none'>A</td>
  <td class=xl66 style='border-top:none;border-left:none'>B</td>
  <td class=xl66 style='border-top:none;border-left:none'>A</td>
  <td class=xl66 style='border-top:none;border-left:none'>B</td>
  <td class=xl66 style='border-top:none;border-left:none'>A</td>
  <td class=xl66 style='border-top:none;border-left:none'>B</td>
  <td class=xl66 style='border-top:none;border-left:none'>A</td>
 </tr>
<!--EndFragment-->
</table>

</body>

</html>
`
    var jsonParams = {
            method: 1,
            params: {
                name: "test测试",
                age: 20,
                attach: {
                    text: text
                }
            }
        },
        jsonStr = JSON.stringify(jsonParams);

    $.ajax({
        url: "http://127.0.0.1:9000/test",
        type: "POST",
        contentType: 'application/x-www-form-urlencoded; charset=UTF-8',
        // contentType: 'application/json',
        // contentType: 'application/x-www-form-urlencoded',
        data: {
            // jsonParams: jsonStr
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

