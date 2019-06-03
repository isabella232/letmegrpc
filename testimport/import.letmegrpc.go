// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: import.proto

package testimport

import (
	encoding_json "encoding/json"
	fmt "fmt"
	serve "github.com/gogo/letmegrpc/letmetestserver/serve"
	proto "github.com/gogo/protobuf/proto"
	golang_org_x_net_context "golang.org/x/net/context"
	google_golang_org_grpc "google.golang.org/grpc"
	io "io"
	log "log"
	math "math"
	net_http "net/http"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var DefaultHtmlStringer = func(req, resp interface{}) ([]byte, error) {
	header := []byte("<p><div class=\"container\"><pre>")
	data, err := encoding_json.MarshalIndent(resp, "", "\t")
	if err != nil {
		return nil, err
	}
	footer := []byte("</pre></div></p>")
	return append(append(header, data...), footer...), nil
}

func Serve(httpAddr, grpcAddr string, stringer func(req, resp interface{}) ([]byte, error), opts ...google_golang_org_grpc.DialOption) {
	handler, err := NewHandler(grpcAddr, stringer, opts...)
	if err != nil {
		log.Fatalf("NewHandler(%q) = %v", grpcAddr, err)
	}
	if err := net_http.ListenAndServe(httpAddr, handler); err != nil {
		log.Fatal(err)
	}
}
func NewHandler(grpcAddr string, stringer func(req, resp interface{}) ([]byte, error), opts ...google_golang_org_grpc.DialOption) (net_http.Handler, error) {
	conn, err := google_golang_org_grpc.Dial(grpcAddr, opts...)
	if err != nil {
		return nil, err
	}
	mux := net_http.NewServeMux()
	OtherLabelClient := NewOtherLabelClient(conn)
	OtherLabelServer := NewHTMLOtherLabelServer(OtherLabelClient, stringer)
	mux.HandleFunc("/OtherLabel/Produce", OtherLabelServer.Produce)
	return mux, nil
}

type htmlOtherLabel struct {
	client   OtherLabelClient
	stringer func(req, resp interface{}) ([]byte, error)
}

func NewHTMLOtherLabelServer(client OtherLabelClient, stringer func(req, resp interface{}) ([]byte, error)) *htmlOtherLabel {
	return &htmlOtherLabel{client, stringer}
}

var FormOtherLabel_Produce string = `<div class="container"><div class="jumbotron">
	<h3>OtherLabel: Produce</h3>
	
	<form class="form-horizontal">
	<div id="form"><div class="children"></div></div>
    <input type="submit" class="btn btn-primary" role="button" value="Submit">
    </form>
    
	<script>

function addChildNode(ev) {
	ev.preventDefault();
	var thisNode = $(this).parents(".node:first");
	var myType = $(this).attr("type");
	var child = $(nodeFactory[myType]);
	activateLinks(child);
	$(">.children[type=" + myType + "]", thisNode).append(child);
}

function setChildNode(ev) {
	ev.preventDefault();
	var thisNode = $(this).parents(".node:first");
	var myType = $(this).attr("type");
	var child = $(nodeFactory[myType]);
	activateLinks(child);
	$(">.children[type=" + myType + "]", thisNode).append(child);
$(">.tooltipper", thisNode).hide();
	$(this).hide();
}

function delChildNode(ev) {
	ev.preventDefault();
	var thisNode = $(this).parents(".node:first");
	var parentNode = thisNode.parents(".node:first");
	thisNode.remove();
	var setChildLink = $(">a.set-child[fieldname='" + thisNode.attr('fieldname') + "']", parentNode);
	if (setChildLink.length > 0) {
		setChildLink.show();
                $(">.tooltipper", parentNode).show();
	}
}

function delField(ev) {
	ev.preventDefault();
	var thisField = $(this).parents(".field:first");
	thisField.remove();
}

function addElem(ev) {
	ev.preventDefault();
	var thisNode = $(this).parents(".node:first");
	var myType = $(this).attr("type");
	var myFieldname = $(this).attr("fieldname");
	if (myType == "bool") {
		var input = $('<div class="field form-group"><label class="col-sm-2 control-label">' + myFieldname + ': </label><div class="col-sm-8"><input name="' + myFieldname + '" type="checkbox" repeated="true"/></div><div class="col-sm-2"><a href="#" class="del-field btn btn-warning btn-sm" role="button">Remove</a></div></div>');
		$("a.del-field", input).click(delField);
		$("> .fields[fieldname='" + myFieldname + "']", thisNode).append(input);
	}
	if (myType == "number") {
		var input = $('<div class="field form-group"><label class="col-sm-2 control-label">' + myFieldname + ': </label><div class="col-sm-8"><input class="form-control" name="' + myFieldname + '" type="number" step="1" repeated="true"/></div><div class="col-sm-2"><a href="#"  class="del-field btn btn-warning btn-sm" role="button">Remove</a></div></div>');
		$("a.del-field", input).click(delField);
		$("> .fields[fieldname='" + myFieldname + "']", thisNode).append(input);
	}
	if (myType == "text") {
		var input = $('<div class="field form-group"><label class="col-sm-2 control-label">' + myFieldname + ': </label><div class="col-sm-8"><input class="form-control" name="' + myFieldname + '" type="text" repeated="true"/></div><div class="col-sm-2"><a href="#"  class="del-field btn btn-warning btn-sm" role="button">Remove</a></div></div>');
		$("a.del-field", input).click(delField);
		$("> .fields[fieldname='" + myFieldname + "']", thisNode).append(input);
	}
	if (myType == "float") {
		var input = $('<div class="field form-group"><label class="col-sm-2 control-label">' + myFieldname + ': </label><div class="col-sm-8"><input class="form-control" name="' + myFieldname + '" type="number" step="any" repeated="true"/></div><div class="col-sm-2"><a href="#"  class="del-field btn btn-warning btn-sm" role="button">Remove</a></div></div>');
		$("a.del-field", input).click(delField);
		$("> .fields[fieldname='" + myFieldname + "']", thisNode).append(input);
	}
}

function getUrlParameter(sParam)
{
    var sPageURL = window.location.search.substring(1);
    var sURLVariables = sPageURL.split('&');
    for (var i = 0; i < sURLVariables.length; i++)
    {
        var sParameterName = sURLVariables[i].split('=');
        if (sParameterName[0] == sParam)
        {
            return sParameterName[1];
        }
    }
}

function activateLinks(node) {
 	$("a.add-child", node).click(addChildNode);
	$("a.set-child", node).click(setChildNode);
	$("a.add-elem", node).click(addElem);
	$("a.del-child", node).click(delChildNode);
	$("a.del-field", node).click(delField);
	$('label[type=checkbox]').click(function() {
	    if ($(this).hasClass('active')) {
	        $(this).removeClass('active');
	    } else {
	        $(this).addClass('active');
	    }
	});
	$('[data-toggle="tooltip"]', node).tooltip();
}

function getChildren(el) {
	var json = {};
	$("> .children > .node", el).each(function(idx, node) {
		var nodeJson = getFields($(node));
		var allChildren = getChildren($(node));
		for (childType in allChildren) {
			nodeJson[childType] = allChildren[childType];
		}
		var nodeType = $(node).attr("fieldname");
		var isRepeated = $(node).attr("repeated") == "true";
		if (isRepeated) {
			if (!(nodeType in json)) {
				json[nodeType] = [];
			}
			json[nodeType].push(nodeJson);
		} else {
			json[nodeType] = nodeJson;
		}
	});
	return json
}

function isInt(value) {
  return !isNaN(value) &&
         parseInt(Number(value)) == value &&
         !isNaN(parseInt(value, 10));
}

function replaceAll(str, search, replace) {
	return str.split(search).join(replace);
}

function escapeIllegal(str) {
	return replaceAll(replaceAll(replaceAll(str, "%", "%25"), "&", "%26"), "#", "%23");
}

function getFields(node) {
	var nodeJson = {};
	$("> div.field > div ", $(node)).each(function(idx, field) {
		$("> input[type=text]", $(field)).each(function(idx, input) {
			nodeJson[$(input).attr("name")] = escapeIllegal($(input).val());
		});
		$("> input[type=number][step=any]", $(field)).each(function(idx, input) {
			nodeJson[$(input).attr("name")] = parseFloat($(input).val());
		});
		$("> input[type=number][step=1]", $(field)).each(function(idx, input) {
			nodeJson[$(input).attr("name")] = parseInt($(input).val());
		});
		$("> div > label.active", $(field)).each(function(idx, label) {
                        var input = $("> input[type=radio]", $(label));
			var v = input.val();
			if (v == "true") {
				nodeJson[input.attr("name")] = true;
			} else if (v == "false") {
				nodeJson[input.attr("name")] = false;
			} else {
				nodeJson[input.attr("name")] = parseInt(input.val());
			}
		});
		$("> select", $(field)).each(function(idx, input) {
			var textvalue = $(input).val();
			if (isInt(textvalue)) {
				nodeJson[$(input).attr("name")] = parseInt(textvalue);
			} else {
				nodeJson[$(input).attr("name")] = escapeIllegal(textvalue);
			}
		});
	});
	$("> div.fields > div ", $(node)).each(function(idx, field) {
		$("input[type=text]", $(field)).each(function(idx, input) {
			var fieldname = $(input).attr("name");
			if (!(fieldname in nodeJson)) {
				nodeJson[fieldname] = [];
			}
			nodeJson[fieldname].push(escapeIllegal($(input).val()));
		});
		$("input[type=checkbox]", $(field)).each(function(idx, input) {
			var fieldname = $(input).attr("name");
			if (!(fieldname in nodeJson)) {
				nodeJson[fieldname] = [];
			}
			nodeJson[fieldname].push($(input).is(':checked'));
		});
		$("input[type=number][step=any]", $(field)).each(function(idx, input) {
			var fieldname = $(input).attr("name");
			if (!(fieldname in nodeJson)) {
				nodeJson[fieldname] = [];
			}
			nodeJson[fieldname].push(parseFloat($(input).val()));
		});
		$("input[type=number][step=1]", $(field)).each(function(idx, input) {
			var fieldname = $(input).attr("name");
			if (!(fieldname in nodeJson)) {
				nodeJson[fieldname] = [];
			}
			nodeJson[fieldname].push(parseInt($(input).val()));
		});
		$("label.active", $(field)).each(function(idx, label) {
                        var input = $("> input[type=radio]", $(label));
			var fieldname = $(input).attr("name");
			if (!(fieldname in nodeJson)) {
				nodeJson[fieldname] = [];
			}
			nodeJson[fieldname].push(parseInt(input.val()));
		});
		$("select", $(field)).each(function(idx, input) {
			var fieldname = $(input).attr("name");
			if (!(fieldname in nodeJson)) {
				nodeJson[fieldname] = [];
			}
			var textvalue = $(input).val();
			if (isInt(textvalue)) {
				nodeJson[fieldname].push(parseInt(textvalue));
			} else {
				nodeJson[fieldname].push(escapeIllegal(textvalue));
			}
		});
	});

	return nodeJson;
}

function radioed(def, index, value) {
	if (value == undefined) {
		if (def == index) {
			return "checked"
		}
		return ""
	}
	if (index == parseInt(value)) {
		return "checked"
	}
	if (index == value) {
		return "checked"
	}
	return ""
}

function activeradio(def, index, value) {
	if (value == undefined) {
		if (def == index) {
			return "active"
		}
		return ""
	}
	if (index == parseInt(value)) {
		return "active"
	}
	if (index == value) {
		return "active"
	}
	return ""
}

function checked(value) {
	if (value == undefined) {
		return ""
	}
	if (value == true) {
		return "checked='checked'"
	}
	return ""
}

function selected(def, index, value) {
	if (value == undefined) {
		if (def == index) {
			return "selected='selected'"
		}
		return ""
	}
	if (index == parseInt(value)) {
		return "selected='selected'"
	}
	if (index == value) {
		return "selected='selected'"
	}
	return ""
}

function emptyIfNull(json) {
	if (json == undefined || json == null) {
		return JSON.parse("{}");
	}
	return json;
}

function getValue(json, name) {
	var value = json[name];
	if (value == undefined) {
		return JSON.parse("{}");
	}
	return value;
}

function getList(json, name) {
	var value = json[name];
	if (value == undefined) {
		return JSON.parse("[]");
	}
	return value;
}

function setLink(json, typ, fieldname, help) {
var display = "";
	if (json[fieldname] != undefined) {
display = 'style="display:none"';
}
        var tooltip = "";
        if (help.length > 0) {
		tooltip = '<a href="#" data-toggle="tooltip" ' + display + ' title="' + help + '" class="tooltipper"><span class="glyphicon glyphicon-question-sign" aria-hidden="true"></span></a>';
        }
	if (json[fieldname] == undefined) {
		return '<a href="#" type="' + typ + '" class="set-child btn btn-success btn-sm" role="button" fieldname="' + fieldname + '">Set ' + fieldname + '</a>' + tooltip;
	}
	return '<a href="#" type="' + typ + '" class="set-child btn btn-success btn-sm" role="button" fieldname="' + fieldname + '" style="display: none;">Set ' + fieldname + '</a>';
}

function setValue(def, value) {
	if (value == undefined) {
		if (def.length == 0) {
			return ""
		}
		return 'value="' + def + '"'
	}
	return 'value="' + value + '"'
}

function setRepValue(value) {
	if (value == undefined) {
		return ""
	}
	return 'value="' + value + '"'
}

function encode_utf8(s) {
  return unescape(encodeURIComponent(s));
}

function decode_utf8(s) {
  return decodeURIComponent(escape(s));
}

function HTMLEncode(str){
  var i = str.length,
      aRet = [];

  while (i--) {
    var iC = str[i].charCodeAt();
    if (iC < 65 || iC > 127 || (iC>90 && iC<97)) {
      aRet[i] = '&#'+iC+';';
    } else {
      aRet[i] = str[i];
    }
   }
  return aRet.join('');
}


function setStrValue(def, value) {
	if (value == undefined) {
		if (def == undefined) {
			return ""
		}
		return "value=" + JSON.stringify(HTMLEncode(decode_utf8(def)));
	}
	return "value=" + JSON.stringify(HTMLEncode(decode_utf8(value)));
}

function setRepStrValue(value) {
	if (value == undefined) {
		return ""
	}
	return "value=" + JSON.stringify(HTMLEncode(decode_utf8(value)));
}

var nodeFactory = {"Album_RootKeyword": buildAlbum_RootKeyword(emptyIfNull(null)),
"RepeatedKeyword_Song_Song": buildRepeatedKeyword_Song_Song(emptyIfNull(null)),
"RepeatedKeyword_Artist_Composer": buildRepeatedKeyword_Artist_Composer(emptyIfNull(null)),}
	function buildRepeatedKeyword_Artist_Composer(json) {
var s = '<div class="node" type="RepeatedKeyword_Artist_Composer" fieldname="Composer" repeated="true">';
s += '<div class="row"><div class="col-sm-2">'
s += '<a href="#" class="del-child btn btn-danger btn-xs" role="button" fieldname="Composer">Remove</a>'
s += '</div><div class="col-sm-10">'
s += '<label class="heading">Composer</label>'
s += '</div></div>'
s += '<div class="field form-group"><label class="col-sm-2 control-label">Name: </label><div class="col-sm-10"><input class="form-control" name="Name" type="text" '+setStrValue("", json["Name"])+'/></div></div>';
				
s += '<div class="field form-group"><label class="col-sm-2 control-label">Role: </label>';
					s += '<div class="col-sm-10"><div class="btn-group" data-toggle="buttons">';
					s += 	'<label class="btn btn-primary ' + activeradio(0, 0, json["Role"]) + '"><input type="radio" name="Role" value="0" ' + radioed(0, 0, json["Role"]) + '/> Voice</label>';
						s += 	'<label class="btn btn-primary ' + activeradio(0, 1, json["Role"]) + '"><input type="radio" name="Role" value="1" ' + radioed(0, 1, json["Role"]) + '/> Guitar</label>';
						s += 	'<label class="btn btn-primary ' + activeradio(0, 2, json["Role"]) + '"><input type="radio" name="Role" value="2" ' + radioed(0, 2, json["Role"]) + '/> Drum</label>';
						s += '</div></div></div>';
					

		s += '</div>';
		return s;
		}

function buildRepeatedKeyword_Song_Song(json) {
var s = '<div class="node" type="RepeatedKeyword_Song_Song" fieldname="Song" repeated="true">';
s += '<div class="row"><div class="col-sm-2">'
s += '<a href="#" class="del-child btn btn-danger btn-xs" role="button" fieldname="Song">Remove</a>'
s += '</div><div class="col-sm-10">'
s += '<label class="heading">Song</label>'
s += '</div></div>'
s += '<div class="field form-group"><label class="col-sm-2 control-label">Name: </label><div class="col-sm-10"><input class="form-control" name="Name" type="text" '+setStrValue("", json["Name"])+'/></div></div>';
				
s += '<div class="field form-group"><label class="col-sm-2 control-label">Track: </label><div class="col-sm-10"><input class="form-control" name="Track" type="number" step="1" '+setValue(0, json["Track"])+'/></div></div>';
				
s += '<div class="field form-group"><label class="col-sm-2 control-label">Duration: </label><div class="col-sm-10"><input class="form-control" name="Duration" type="number" step="any" '+setValue(0, json["Duration"])+'/></div></div>';
				
s += '<div class="children" type="RepeatedKeyword_Artist_Composer">';
			var Composer = getList(json, "Composer");
			for (var i = 0; i < Composer.length; i++) {
				s += buildRepeatedKeyword_Artist_Composer(Composer[i]);
			}
			s += '</div>';
			s += '<a href="#" class="add-child btn btn-success btn-sm" role="button" type="RepeatedKeyword_Artist_Composer">add Composer</a>';
			s += '<div class="field form-group"></div>';
			

		s += '</div>';
		return s;
		}

function buildAlbum_RootKeyword(json) {
if (json == undefined) {
		return "";
	}
	
var s = '<div class="node" type="Album_RootKeyword" fieldname="RootKeyword" repeated="false">';
s += '<div class="field form-group"><label class="col-sm-2 control-label">Name: </label><div class="col-sm-10"><input class="form-control" name="Name" type="text" '+setStrValue("", json["Name"])+'/></div></div>';
				
s += '<div class="children" type="RepeatedKeyword_Song_Song">';
			var Song = getList(json, "Song");
			for (var i = 0; i < Song.length; i++) {
				s += buildRepeatedKeyword_Song_Song(Song[i]);
			}
			s += '</div>';
			s += '<a href="#" class="add-child btn btn-success btn-sm" role="button" type="RepeatedKeyword_Song_Song">add Song</a>';
			s += '<div class="field form-group"></div>';
			
s += '<div class="field form-group"><label class="col-sm-2 control-label">Genre: </label><div class="col-sm-10">';
					s += '<select class="form-control" name="Genre">';
					s += 	'<option value="0" ' + selected(0, 0, json["Genre"]) + '>Pop</option>';
						s += 	'<option value="1" ' + selected(0, 1, json["Genre"]) + '>Rock</option>';
						s += 	'<option value="2" ' + selected(0, 2, json["Genre"]) + '>Jazz</option>';
						s += 	'<option value="3" ' + selected(0, 3, json["Genre"]) + '>NintendoCore</option>';
						s += 	'<option value="4" ' + selected(0, 4, json["Genre"]) + '>Indie</option>';
						s += 	'<option value="5" ' + selected(0, 5, json["Genre"]) + '>Punk</option>';
						s += 	'<option value="6" ' + selected(0, 6, json["Genre"]) + '>Dance</option>';
						s += '</select></div></div>';
					
s += '<div class="field form-group"><label class="col-sm-2 control-label">Year: </label><div class="col-sm-10"><input class="form-control" name="Year" type="text" '+setStrValue("", json["Year"])+'/></div></div>';
				
s += '<div class="fields" fieldname="Producer">';
				var Producer = getList(json, "Producer");
				for (var i = 0; i < Producer.length; i++) {
					s += '<div class="field form-group"><label class="col-sm-2 control-label">Producer: </label><div class="col-sm-8"><input class="form-control" name="Producer" type="text" repeated="true" '+setRepStrValue(json["Producer"][i])+'/></div><div class="col-sm-2"><a href="#" class="del-field btn btn-warning btn-sm" role="button">Remove</a></div></div>';
				}
				s += '</div>';
				s += '<a href="#" fieldname="Producer" class="add-elem btn btn-info btn-sm" role="button" type="text">add Producer</a>';
				s += '<div class="field form-group"></div>';
				
s += '<div class="field form-group"><label class="col-sm-2 control-label">Mediocre: </label>';
					s += '<div class="col-sm-10"><div class="btn-group" data-toggle="buttons">';
					s += 	'<label class="btn btn-primary ' + activeradio(false, false, json["Mediocre"]) + '"><input type="radio" name="Mediocre" value="false" ' + radioed(false, false, json["Mediocre"]) + '/>No</label>';
					s += 	'<label class="btn btn-primary ' + activeradio(false, true, json["Mediocre"]) + '"><input type="radio" name="Mediocre" value="true" ' + radioed(false, true, json["Mediocre"]) + '/>Yes</label>';
					s += '</div></div></div>';
					
s += '<div class="field form-group"><label class="col-sm-2 control-label">Rated: </label>';
					s += '<div class="col-sm-10"><div class="btn-group" data-toggle="buttons">';
					s += 	'<label class="btn btn-primary ' + activeradio(false, false, json["Rated"]) + '"><input type="radio" name="Rated" value="false" ' + radioed(false, false, json["Rated"]) + '/>No</label>';
					s += 	'<label class="btn btn-primary ' + activeradio(false, true, json["Rated"]) + '"><input type="radio" name="Rated" value="true" ' + radioed(false, true, json["Rated"]) + '/>Yes</label>';
					s += '</div></div></div>';
					
s += '<div class="field form-group"><label class="col-sm-2 control-label">Epilogue: </label><div class="col-sm-10"><input class="form-control" name="Epilogue" type="text" '+setStrValue("", json["Epilogue"])+'/></div></div>';
				

				s += '<div class="fields" fieldname="Likes">';
				var Likes = getList(json, "Likes");
				for (var i = 0; i < Likes.length; i++) {
					s += '<div class="field form-group"><label class="col-sm-2 control-label">Likes: </label><div class="col-sm-8"><input name="Likes" type="checkbox" repeated="true" ' + checked(Likes[i]) + '/></div><div class="col-sm-2"><a href="#" class="del-field btn btn-warning btn-sm" role="button">Remove</a></div></div>';
				}
				s += '</div>';
				s += '<a href="#" fieldname="Likes" class="add-elem btn btn-info btn-sm" role="button" type="bool">add Likes</a>';
				s += '<div class="field form-group"></div>';
				

			s += '</div>';
			var node = $(s);
			activateLinks(node);
			return node;
		}function init() {
	var root = $(nodeFactory["Album_RootKeyword"]);
	var jsonText = getUrlParameter("json");
	if (jsonText == undefined) {
		var json = emptyIfNull(null);
	} else {
		var json = JSON.parse(unescape(jsonText));
	}
	$("#form > .children").html(buildAlbum_RootKeyword(json));
	activateLinks(root);
	$("form").submit(function(ev) {
		ev.preventDefault();
		c = getChildren($("#form"));
		j = JSON.stringify(c["RootKeyword"]);
		window.location.assign("./Produce?json="+j);
	});
}

	init();

	</script>

	<style>

	.node{
		padding-left: 2em;
		min-height:20px;
	    padding:10px;
	    margin-top:10px;
	    margin-bottom:20px;
	    //border-left:0.5px solid #999;
	    -webkit-border-radius:4px;
	    -moz-border-radius:4px;
	    border-radius:4px;
	    -webkit-box-shadow:inset 0 1px 1px rgba(0, 0, 0, 0.05);
	    -moz-box-shadow:inset 0 1px 1px rgba(0, 0, 0, 0.05);
	    box-shadow:inset 0 1px 1px rgba(0, 0, 0, 0.05);
	    background-color:#eaeaea;
	}

	.node .node {
		background-color:#e2e2e2;
	}

	.node .node .node {
		background-color:#d9d9d9;
	}

	.node .node .node .node {
		background-color:#d1d1d1;
	}

	.node .node .node .node .node {
		background-color:#c7c7c7;
	}

	.node .node .node .node .node .node {
		background-color:#c0c0c0;
	}

	label{
	        font-weight: normal;
	}

	.heading {
		font-weight: bold;
	}

	</style>
	
	</div>`

func (this *htmlOtherLabel) Produce(w net_http.ResponseWriter, req *net_http.Request) {
	w.Write([]byte(Header(`OtherLabel`, `Produce`)))
	jsonString := req.FormValue("json")
	someValue := false
	msg := &serve.Album{}
	if len(jsonString) > 0 {
		err := encoding_json.Unmarshal([]byte(jsonString), msg)
		if err != nil {
			if err != io.EOF {
				w.Write([]byte("<div class=\"alert alert-danger\" role=\"alert\">" + err.Error() + "</div>"))
				return
			}
			w.Write([]byte("<div class=\"alert alert-danger\" role=\"alert\">" + err.Error() + "</div>"))
		}
		someValue = true
	}
	w.Write([]byte(FormOtherLabel_Produce))
	if someValue {
		reply, err := this.client.Produce(golang_org_x_net_context.Background(), msg)
		if err != nil {
			if err != io.EOF {
				w.Write([]byte("<div class=\"alert alert-danger\" role=\"alert\">" + err.Error() + "</div>"))
				return
			}
			w.Write([]byte("<div class=\"alert alert-danger\" role=\"alert\">" + err.Error() + "</div>"))
		}
		out, err := this.stringer(msg, reply)
		if err != nil {
			if err != io.EOF {
				w.Write([]byte("<div class=\"alert alert-danger\" role=\"alert\">" + err.Error() + "</div>"))
				return
			}
			w.Write([]byte("<div class=\"alert alert-danger\" role=\"alert\">" + err.Error() + "</div>"))
		}
		w.Write(out)
	}
	w.Write([]byte(Footer))
}

var Header func(servName, methodName string) string = func(servName, methodName string) string {
	return `
	<html>
	<head>
	<title>` + servName + `:` + methodName + `</title>
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap.min.css">
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/js/bootstrap.min.js"></script>
	</head>
	<body>
	`
}
var Footer string = `
	</body>
	</html>
	`
