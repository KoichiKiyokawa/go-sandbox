-- Create "users" table
CREATE TABLE `users` (`id` text NULL, `name` text NULL, PRIMARY KEY (`id`));
-- Create "posts" table
CREATE TABLE `posts` (`id` text NULL, `user_id` text NULL, `title` text NULL, `body` text NULL, PRIMARY KEY (`id`), CONSTRAINT `0` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION);
