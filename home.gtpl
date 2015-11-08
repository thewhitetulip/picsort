<html>
	<head> 
		<title> Pic sorter {{ if ne .Remaining 0 }} {{.Name}} {{else}} you are done {{end}}</title>
	</head>

	<body>
		
		{{if eq .Remaining 0}}<p>That is the end of it, <a href="/sort/" >sort</a> them now</p> 
		
		{{ else }} 
		
		<p>{{.Remaining}} picture(s) are left</p>
		
		<img src="/static/{{.Name}}" height="500" width="900">
		<form action="/" method="POST">
			<input type="text" name="tags" />
			<input type="text" name="photoname" type="hidden" value="{{.Name}}"/>
			<input type="submit" name="submit" />
		</form>
		{{end}}
	</body>

</html>
