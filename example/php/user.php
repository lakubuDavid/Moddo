<?php
class User {
	public function __construct(int $id, string $name, string $email, string $posts, string $password, string $birthdate) {
		$this->id = $id;
		$this->name = $name;
		$this->email = $email;
		$this->posts = $posts;
		$this->password = $password;
		$this->birthdate = $birthdate;
	}
	public $id; 
	public $name; 
	public $email; 
	public $posts; 
	public $password; 
	public $birthdate; 
}
?>
