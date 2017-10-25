NAME = $(notdir $(PWD))

build:
	go build -o app
	docker build -t $(NAME):latest -f Dockerfile .

push@%:
	$(eval VERSION = latest)
	$(eval TAG = $*/$(NAME):$(VERSION))
	docker tag $(NAME):$(VERSION) $(TAG)
	docker push $(TAG)
