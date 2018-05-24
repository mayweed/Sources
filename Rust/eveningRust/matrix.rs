use ipoint2::*;

#[derive(Clone,Debug)]
pub struct Matrix<T> {
	cells : Box<[T]>,
	pub width : i32,
	pub height : i32,
}

impl<T : Copy + Default> Matrix<T> {
	pub fn new( x : i32, y : i32 ) -> Matrix<T> {
		Matrix {
			cells: vec![T::default(); (x*y) as usize].into_boxed_slice(),
			width: x,
			height: y,
		}
	}

	pub fn resize_and_fill(&mut self, x : i32, y : i32, v : T ) {
		self.width = x;
		self.height = y;
		let len = (self.width*self.height) as usize;
		if len > self.cells.len() {
			self.cells = vec![v; len].into_boxed_slice();
		} else {
			self.fill(v);
		}
	}
	
	pub fn fill(&mut self, v : T ) {
		for i in 0..(self.linear_len()) {
			self.cells[i as usize] = v;
		}
	}
	
	pub fn linear_len( &self ) -> i32 {
		self.width * self.height
	}
	
	pub fn linear_get( &self, loc : i32 ) -> T {
		self.cells[loc as usize]
	}
	
	pub fn linear_set( &mut self, loc : i32, v : T ) {
		self.cells[loc as usize] = v;
	}
	
	pub fn get( &self, loc : IPoint2 ) -> T {
		self.cells[ (loc.y*self.width + loc.x) as usize ]
	}
	pub fn get_xy( &self, x : i32, y : i32 ) -> T {
		self.cells[ (y*self.width + x) as usize ]
	}

	pub fn get_def( &self, loc : IPoint2, def : T ) -> T {
		if self.is_valid( loc ) { self.get(loc ) } else { def }
	}
	
	pub fn set( &mut self, loc : IPoint2, v : T ) {
		self.cells[ (loc.y*self.width + loc.x) as usize ] = v;
	}
	
	pub fn is_valid( &self, loc : IPoint2 ) -> bool {
		return loc.x >= 0 && loc.x < self.width && 
			loc.y >= 0 && loc.y < self.height;
	}
}
