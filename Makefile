all: day_1
.PHONY: aoc
day_1:
	@docker build . --target bin \
	--output ./day_1/out/