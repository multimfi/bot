package http

var indexHTML = []byte(`<!DOCTYPE html>
<html lang="en">
<head>
<title>bot</title>
<script type="text/javascript">
window.onload = function () {
	var conn;

	var interval = 0;
	var alerts = {};
	var responders = {};

	var state = document.getElementById("state").tBodies[0];
	var states = ["firing", "resolved"];
	var labels = {
		"labels": function(obj, item) {
			for (x in obj["labels"]) {
				if (x == "responders")
					continue
				item.appendChild(kv(x, obj["labels"][x]));
			}
		},

		"annotations": function(obj, item) {
			for (x in obj["annotations"]) {
				item.appendChild(kv(x, obj["annotations"][x]));
			}
		},

		"startsAt": function(obj, item) {
			if (obj.status != "firing") {
				delete(alerts[obj.hash].updater)
				return;
			}
			alerts[obj.hash].updater = function(t) {
				var d = new Date(obj["startsAt"]);
				item.innerHTML = obj["startsAt"]
				item.innerHTML += "<br><b>" + since(t, d) + "</b>";
			}
			alerts[obj.hash].updater(Date.now())
		},

		"endsAt": function(obj, item) {
			var d1 = new Date(obj["endsAt"])
			var d2 = new Date(obj["startsAt"])
			if (d1 < 1) {
				item.innerHTML = "tbd";
				return
			}
			item.innerHTML = obj["endsAt"]
			item.innerHTML += "<br><b>"+since(d1, d2)+"</b>";
		},

		"responders": function(obj, item) {
			if (obj.status != "firing") {
				return;
			}

			var r = obj["responders"];
			if (!r) {
				item.innerHTML = "none";
				return;
			}
			item.innerHTML = "<b>" + r[obj.current] + "</b>";
			item.innerHTML += "<br>" + obj.responders;
		}
	}

	function updater() {
		var now = Date.now();
		for (i in alerts) {
			if (typeof alerts[i].updater != "undefined") {
				alerts[i].updater(now);
			}
		}
	}

	function since(x, y) {
		var t = (x - y)/1e3;
		var s = Math.floor(t%60);
		t /= 60
		var m = Math.floor(t%60);
		t /= 60
		var h = Math.floor(t)
		return h+"h"+m+"m"+s+"s";
	}

	function kv(k, v) {
		var ret = document.createElement("p");
		var b = document.createElement("b");
		b.innerHTML = k;
		ret.appendChild(b);
		ret.innerHTML += " " + v;
		return ret;
	}

	function setState(item, state) {
		for (x in states) {
			if (state != states[x]) {
				item.classList.remove(states[x]);
				continue;
			}
			item.classList.add(state);
		}
	}

	function getChild(name, obj) {
		for (x in obj.children) {
			y = obj.children[x]
			if (y.name != name)
				continue
			return y
		}
	}

	function updateState(obj) {
		var alertItem;
		if (!alerts[obj.hash]) {
			alertItem = document.createElement("tr");
			alertItem.classList.add("alert")
			alerts[obj.hash] = alertItem;
			state.appendChild(alertItem);
		} else {
			alertItem = alerts[obj.hash];
		}

		setState(alertItem, obj.status);
		alertItem.id = obj.hash;
		for (name in labels) {
			var item = getChild(name, alertItem);
			if (!item) {
				item = document.createElement("td");
				alertItem.appendChild(item)
				item.name = name;
				labels[name](obj, item)
			}
			switch(name) {
			case "endsAt":
			case "startsAt":
			case "responders":
				labels[name](obj, item);
			}
		}
	}

	function updateIRC(obj) {
		var e = document.getElementById("irc")
		if (obj.status != true) {
			e.style.color = "red";
			e.innerHTML = "FAIL!";
			return
		}
		e.style.color = "green";
		e.innerHTML = "OK!";
	}

	colormap = {
		"active": "green",
		"failed": "red",
		"unknown": "black",
	}

	function updateResponders(obj) {
		var e = document.getElementById("responders")
		for (r in obj) {
			var res = obj[r];
			var item = getChild(res.name, e);
			if (!item) {
				item = document.createElement("span");
				e.appendChild(item)
				item.name = res.name;
				item.innerHTML = res.name;
			}
			item.style.color = colormap[res.state];
		}
	}

	if (!window["WebSocket"]) {
		var item = document.createElement("div");
		item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
		appendLog(item);
		return;
	}

	conn = new WebSocket("ws://" + document.location.host + "/ws");
	conn.onopen = function() {
		conn.send("alerts");
		conn.send("irc");
		conn.send("responders");
	};
	conn.onclose = function (evt) {
		var item = document.createElement("div");
		item.innerHTML = "<b>Connection closed.</b>";
		document.body.append(item);
	};
	conn.onmessage = function (evt) {
		var obj = JSON.parse(evt.data);
		switch(obj.type) {
		case "alert":
			updateState(obj.msg);
			break;
		case "irc":
			updateIRC(obj.msg);
			break;
		case "responders":
			updateResponders(obj.msg);
			break;
		}

		if (document.getElementsByClassName("firing").length < 1) {
			clearInterval(interval);
			interval = 0;
			return;
		}

		if (interval > 0) {
			return;
		}

		interval = setInterval(updater, 1000);
	};
};
</script>
<style type="text/css">
body {
	font-family: sans-serif;
}

td p {
	margin: 0;
	padding: 0;
}

table tr th {
	text-align: left;
}

table {
	width: 100%;
	border-spacing: 0;
}

table tr td, table tr th, h1 {
	padding-left: 1em;
}

.firing {
	background: rgb(255, 156, 156);
}

.resolved {
	background: rgb(232, 255, 240);
}

td {
	padding: 1em;
	border-top: 5px solid #fff;
}

#responders span {
	font-weight: bold;
	padding: 1em;
}
</style>
</head>
<body>
<p>IRC <b><span id="irc">TBD</span></b></p>
<h4>Responders</h4>
<div id="responders"></div>
<h3>Alerts</h3>
<div>
<table id="state">
  <tr>
    <th>Labels</th> 
    <th>Annotations</th>
    <th>Start</th>
    <th>End</th>
    <th>Responder</th>
  </tr>
</table>
</div>
</body>
</html>
`)

// vim: set ft=javascript:
