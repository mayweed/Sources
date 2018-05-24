use std::ops::*;

#[derive(Default, Clone, Copy, Debug)]
pub struct Point2 {
	pub x: f64,
	pub y: f64,
}

impl Point2 {
	pub fn at( x : f64, y : f64 ) -> Self {
		Point2 { x, y }
	}

	pub fn sub(&self, o: Point2) -> Self {
		Point2 {
			x: self.x - o.x,
			y: self.y - o.y,
		}
	}

	pub fn len(&self) -> f64 {
		(self.x * self.x + self.y * self.y).sqrt()
	}

	pub fn norm(&self) -> Point2 {
		let l = self.len();
		if l < 1e-4f64 {
			return Point2 { x: 0f64, y: 0f64 };
		}
		self.div(l)
	}

	pub fn div(&self, s: f64) -> Point2 {
		Point2 {
			x: self.x / s,
			y: self.y / s,
		}
	}

	pub fn add(&self, o: Point2) -> Point2 {
		Point2 {
			x: self.x + o.x,
			y: self.y + o.y,
		}
	}

	pub fn mul(&self, s: f64) -> Point2 {
		Point2 {
			x: self.x * s,
			y: self.y * s,
		}
	}

	pub fn dist_to(&self, o: Point2) -> f64 {
		self.sub(o).len()
	}

	pub fn factor_to(&self, o : Point2, f : f64 ) -> Point2 {
		*self + (o - *self) * f
	}

	//directional from origin
	pub fn project_onto(&self, o : Point2) -> Point2 {
		let no = o.norm();
		let q = self.dot( no );
		q * no
	}

	//directional from origin
	pub fn reject_onto(&self, o : Point2) -> Point2 {
		*self - self.project_onto(o)
	}


	pub fn dot( &self, o : Point2) -> f64 {
		self.x*o.x + self.y*o.y
	}

	pub fn dist_to_line( &self, pt : Point2, norm : Point2 ) -> f64 {
		(*self - pt).reject_onto( norm ).len()
	}
}

impl Add<Point2> for Point2 {
	type Output = Point2;
	fn add(self, o : Point2) -> Self {
		Point2 {
			x: self.x + o.x,
			y: self.y + o.y,
		}
	}
}

impl Sub<Point2> for Point2 {
	type Output = Point2;
	fn sub(self, o : Point2) -> Self {
		Point2 {
			x: self.x - o.x,
			y: self.y - o.y,
		}
	}
}

impl Div<Point2> for Point2 {
	type Output = Point2;
	fn div(self, o : Point2) -> Self {
		Point2 {
			x: self.x / o.x,
			y: self.y / o.y,
		}
	}
}

impl Div<f64> for Point2 {
	type Output = Point2;
	fn div(self, s : f64) -> Self {
		Point2 {
			x: self.x / s,
			y: self.y / s,
		}
	}
}

impl Mul<Point2> for Point2 {
	type Output = Point2;
	fn mul(self, o : Point2) -> Self {
		Point2 {
			x: self.x * o.x,
			y: self.y * o.y,
		}
	}
}

impl Mul<f64> for Point2 {
	type Output = Point2;
	fn mul(self, s : f64) -> Self {
		Point2 {
			x: self.x * s,
			y: self.y * s,
		}
	}
}

impl Mul<Point2> for f64 {
	type Output = Point2;
	fn mul(self, s : Point2) -> Point2 {
		Point2 {
			x: self * s.x,
			y: self * s.y,
		}
	}
}
