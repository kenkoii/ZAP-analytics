$(document).ready(function(){
  const password = 'd880d38272b0c35cb21fe6a7f3038959';
  $('#generate').click(function(e){
      e.preventDefault();
      console.log("Yoohoo")
      var params = loadParams();
      console.log(params);
      $('#generate').attr({hidden: true});
      $('#loader').css({'display': 'block'});
      if($.md5($('#password').val()) === password) {
        $('#form').submit();
        // $.get("https://frecre-analytics.appspot.com/api/v1/analytics/" + $('#table').val(), params, function(data, status){

        //   var csv = Papa.unparse(data);
        //   var JSONData = JSON.parse(data);
        //   console.log(csv);
        //   console.log(JSONData);
        //   showDownloadCSVButton(csv, $('#table').val(), params);
        //   //  buildHtmlTable(JSON.parse(data));
          
        //   $('#result-div').append("<p style='font-size: 15px'>Number of rows returned: <strong style='color: #000'>" + JSONData.length + "</strong></p>");
        //   $('#result-div').append('<textarea id="result" style="width: 100%" rows="20" hidden readonly></textarea>');
        //   $('#result').show().val(JSON.stringify(JSONData, undefined, 4));
        //   toggleButtons();
        // });
      } else {
        alert("Error: Incorrect Password");
        $('#generate').attr({hidden: null});
        $('#loader').css({'display': 'none'});
      }
      // $.get("http://0f291035.ngrok.io/api/v1/analytics/" + $('#table').val(), params, function(data, status){
      
      // $(this).attr({'hidden': true});
      // $('#reset').attr({})
  });

  $('#table').on('change', function() {
    console.log("Hello!");
    if(this.value == 'userproperties') {
      $('#filter').css({'display': 'block'});
    } else {
      $('#filter').css({'display': 'none'});
    }
  });

  $('#reset').click(function(e){
    resetPage();
  });
});

function toggleButtons() {
  // var generate = $('#generate').attr('hidden');
  // var reset = $('#reset').attr('hidden')
  // $('#generate').attr({hidden: true});
  $('#loader').css({'display': 'none'});
  $('#reset').attr({hidden: null});
  $('input').attr({'disabled': true});
  $('select').prop('disabled', 'disabled');
}

function resetPage() {
  console.log('Hello');
  // window.location.reload(true);
  window.location.href = window.location.href;
}

function showDownloadCSVButton(csv, table, params) {
  // Data URI
  // csvData = 'data:application/csv;charset=utf-8,' + encodeURIComponent(csv);
  xData = new Blob([csv], { type: 'text/csv' });
  var xUrl = URL.createObjectURL(xData);
  // a.href = xUrl;
  const filename = params.start + params.end + table;
  $('#download').css({'display': 'block'})
                .attr({'download': filename, 'href': xUrl, 'target': '_blank'});
}

function hideDownloadButton() {
  $('#download').css({'display': 'none'});
}

function loadParams() {
  var params = {};
  if($('#start').val()) {
    params['start'] = $('#start').val();
  }
  if($('#end').val()) {
    params['end'] = $('#end').val();
  }
  if($("input[name='filter']:checked").val()) {
    // params['filter'] = $('#filter').val();
    params['filter'] = $("input[name='filter']:checked").val();
  }

  return params;
}


function loadFields(params) {
  $('#start').val(params['start']);
  $('#end').val(params['end']);
  $('#table').val(params['table']);
  $('#filter').val(params['filter']);
}

// Builds the HTML Table out of myList.
function buildHtmlTable(myList) {
  var columns = addAllColumnHeaders(myList);

   for (var i = 0 ; i < myList.length ; i++) {
       var row$ = $('<tr/>');
       for (var colIndex = 0 ; colIndex < columns.length ; colIndex++) {
           var cellValue = myList[i][columns[colIndex]];

           if (cellValue == null) { cellValue = ""; }

           row$.append($('<td/>').html(cellValue));
       }
       $("#excelDataTable").append(row$);
   }
}

// Adds a header row to the table and returns the set of columns.
// Need to do union of keys from all records as some records may not contain
// all records.
function addAllColumnHeaders(myList) {
  var columnSet = [];
  var headerTr$ = $('<tr/>');

  for (var i = 0; i < myList.length; i++) {
    var rowHash = myList[i];
    for (var key in rowHash) {
      if ($.inArray(key, columnSet) == -1) {
        columnSet.push(key);
        headerTr$.append($('<th/>').html(key));
      }
    }
  }
  $("#excelDataTable").append(headerTr$);

  return columnSet;
}

function syntaxHighlight(json) {
    json = json.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
    return json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, function (match) {
        var cls = 'number';
        if (/^"/.test(match)) {
            if (/:$/.test(match)) {
                cls = 'key';
            } else {
                cls = 'string';
            }
        } else if (/true|false/.test(match)) {
            cls = 'boolean';
        } else if (/null/.test(match)) {
            cls = 'null';
        }
        return '<span class="' + cls + '">' + match + '</span>';
    });
}
