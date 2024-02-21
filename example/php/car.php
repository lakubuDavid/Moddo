<?php
class Car {
	public function __construct(int $id, string $owner, string $color, mixed $paint, float $fuel) {
		$this->id = $id;
		$this->owner = $owner;
		$this->color = $color;
		$this->paint = $paint;
		$this->fuel = $fuel;
	}
	public $id; 
	public $owner; 
	public $color; 
	public $paint; 
	public $fuel; 
}
?>
