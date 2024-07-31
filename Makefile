

templ:
	templ generate --watch --proxy="http://localhost:6969"

tailwind:
	tailwind --config tailwind.config.js -i config/style.css -o static/css/style.css -w --minify	
