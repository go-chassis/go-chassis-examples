package com.xxx.order;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.apache.servicecomb.springboot.starter.provider.EnableServiceComb;

@SpringBootApplication
@EnableServiceComb
public class Application {

	public static void main(String[] args) {
		SpringApplication.run(Application.class, args);
	}
}
