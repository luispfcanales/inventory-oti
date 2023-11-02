run:
	go build .
css:
	npx tailwindcss -i ./tailwindcss-input.css -o ./public/dist/output.css --watch
