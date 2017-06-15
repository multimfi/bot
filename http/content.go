package http

var indexHTML = []byte(`<!DOCTYPE html>
<html lang="en">
<head>
<title>bot</title>
<script type="text/javascript">
window.onload = function () {
	var conn;
	var alerts = {};
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
			if (obj.status != "firing")
				return;
			var d = new Date(obj["startsAt"]);
			item.innerHTML = since(Date.now(), d);
		},

		"endsAt": function(obj, item) {
			var d1 = new Date(obj["endsAt"])
			var d2 = new Date(obj["startsAt"])
			if (d1 < 1) {
				item.innerHTML = "tbd";
				return
			}
			item.innerHTML = since(d1, d2);
		},

		"responders": function(obj, item) {
			var r = obj["responders"];
			if (!r) {
				item.innerHTML = "none";
				return;
			}
			item.innerHTML = r[obj.current];
		}
	}


	function since(x, y) {
		var t = (x - y)/1e3;
		var h = Math.floor(t/3600);
		var m = Math.floor((t/60)%60);
		var s = Math.floor(t%60);
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
				labels[name](obj, item);
				break;
			case "startsAt":
				labels[name](obj, item);
				break;
			}
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
		conn.send("status");
	};
	conn.onclose = function (evt) {
		var item = document.createElement("div");
		item.innerHTML = "<b>Connection closed.</b>";
		document.body.append(item);
	};

	conn.onmessage = function (evt) {
		var obj = JSON.parse(evt.data);
		updateState(obj);
	};
};
</script>
<style type="text/css">
body, td p {
	font-family: sans-serif;
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

.alert {
	padding: 0.1%;
	border-bottom: 1px solid #ddd;
}

.alert p, .item{
	margin: auto;
}
</style>
</head>
<body>
<div>
<table id="state">
  <tr>
    <th>Labels</th> 
    <th>Annotations</th>
    <th>Since</th>
    <th>Lasted</th>
    <th>Responder</th>
  </tr>
</table>
</div>
</body>
</html>
`)

// vim: set ft=javascript:
