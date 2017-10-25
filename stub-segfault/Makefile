NAME = $(notdir $(PWD))

build:
	gcc -o app segfault.c
	docker build -t $(NAME):latest -f Dockerfile .

push@%:
	$(eval VERSION = latest)
	$(eval TAG = $*/$(NAME):$(VERSION))
	docker tag $(NAME):$(VERSION) $(TAG)
	docker push $(TAG)
