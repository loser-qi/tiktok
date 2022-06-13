create table user(
	id bigint primary key auto_increment,
	username varchar(100) not null,
	password varchar(100) not null,
	follow_count bigint not null default 0,
	follower_count bigint not null default 0
)auto_increment=10001;

create table video(
	id bigint primary key auto_increment,
	user_id bigint not null,
	title varchar(200) not null,
	play_path varchar(200) not null,
	cover_path varchar(200) not null,
	favorite_count bigint not null default 0,
	comment_count bigint not null default 0,
	create_time datetime not null,
	constraint fk_video_user foreign key (user_id) references user(id)
);

create table comment(
	id bigint primary key auto_increment,
	user_id bigint not null,
	video_id bigint not null,
	text text not null,
	create_time datetime not null,
	constraint fk_is_comment_user foreign key (user_id) references user(id),
	constraint fk_is_comment_video foreign key (video_id) references video(id)
);

create table is_favorite(
	user_id bigint not null,
	video_id bigint not null,
	action_type int not null default 1,
	primary key(user_id,video_id),
	constraint fk_is_favorite_user foreign key (user_id) references user(id),
	constraint fk_is_favorite_video foreign key (video_id) references video(id)
);

create table is_follow(
	follower_user_id bigint not null,
	followed_user_id bigint not null,
	action_type int not null default 1,
	primary key(followed_user_id,follower_user_id),
	constraint fk_is_comment_followed_user foreign key (followed_user_id) references user(id),
	constraint fk_is_comment_follower_user foreign key (follower_user_id) references user(id)
);