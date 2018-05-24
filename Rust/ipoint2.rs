//Thx to evening rust for that!!
use std::ops::*;

#[derive(Default,Copy,Clone,Debug,PartialEq)]
pub struct IPoint2 {
	pub x : i32,
	pub y : i32,
}

impl IPoint2 {
	pub fn at( x : i32, y : i32 ) -> Self {
		Self {
			x,
			y,
		}
	}

	pub fn manhatten_dist( &self, to : IPoint2 ) -> i32 {
		(to.x - self.x).abs() + (to.y - self.y).abs()
	}

	pub fn dir_to( &self, to : IPoint2 ) -> IPoint2 {
		let off = to - *self;
		IPoint2 {
			x: off.x.signum(),
			y: off.y.signum(),
		}
	}
}

impl Add<IPoint2> for IPoint2 {
	type Output = IPoint2;
	fn add(self, o : IPoint2) -> Self {
		IPoint2 {
			x: self.x + o.x,
			y: self.y + o.y,
		}
	}
}

impl Sub<IPoint2> for IPoint2 {
	type Output = IPoint2;
	fn sub(self, o : IPoint2) -> Self {
		IPoint2 {
			x: self.x - o.x,
			y: self.y - o.y,
		}
	}
}

impl Mul<i32> for IPoint2 {
	type Output = IPoint2;
	fn mul(self, s : i32) -> Self {
		IPoint2 {
			x: self.x * s,
			y: self.y * s,
		}
	}
}

impl Mul<IPoint2> for i32 {
	type Output = IPoint2;
	fn mul(self, pt : IPoint2) -> IPoint2 {
		IPoint2 {
			x: self * pt.x,
			y: self * pt.y,
		}
	}
}
