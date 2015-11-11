<html>
	<head> 
		<title> Pic sorter {{ if ne .Remaining 0 }} {{.Name}} {{else}} you are done {{end}}</title>
		
		<style>
			img {
				height: 640px;
				width: 950px;
			}
			
			.hidden{
				visibility:hidden;
				width: 0px;
			}
		</style>
	</head>

	<body>
		
		{{if eq .Remaining 0}}
				<p>That is the end of it, <a href="/sort/" >sort</a> them now</p>
		{{ else }} 
		<span style="float:left;">
		<img src="/static/{{.Name}}" >
		</span>
		<span style="float:right;padding-right: 5%;padding-top: 25%;">
			
			<p>{{.Remaining}} picture(s) are left</p>			
			<form action="/" method="POST">
				<input type="text" name="tags" />
				<input type="submit" name="submit" />
				<input type="text" name="photoname" class="hidden" value="{{.Name}}"/>
			</form>
		</span>
	
		{{end}}
	</body>

</html>
