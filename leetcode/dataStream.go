/**
* Definition for an interval.
* type Interval struct {
	*	   Start int
	*	   End   int
	* }
	*/
	type SummaryRanges struct {
		interval []Interval
	}


	/** Initialize your data structure here. */
	func Constructor() SummaryRanges {
		return SummaryRanges{interval: make([]Interval, 0)}
	}


	func (this *SummaryRanges) Addnum(val int)  {
		if len(this.interval) == 0 {
			this.interval = append(this.interval, Interval{val, val})
		}

		if val < this.interval[0].Start - 1 {
			this.interval = append([]Interval{Interval{val, val}}, this.interval...)
			return
		}

		for i := 0; i < len(this.interval); i++ {
			if val >= this.interval[i].Start && val <= this.interval[i].End {
				return
			}
			if val == this.interval[i].Start - 1 {
				this.interval[i].Start = val
				return
			}
			if val == this.interval[i].End + 1 {
				this.interval[i].End = val
				if (i < len(this.interval) - 1) && (this.interval[i+1].Start == val + 1) {
					this.interval[i+1].Start = this.interval[i].Start
					this.interval = append(this.interval[:i], this.interval[i+1:]...)
				}
				return
			}
			if val < this.interval[i].Start {
				next := this.interval[i:]
				this.interval = append(this.interval, Interval{val, val})
				this.interval = append(this.interval[:i+1], next...)
				return
			}
		}

		this.interval = append(this.interval, Interval{val, val})

	}


	func (this *SummaryRanges) Getintervals() []Interval {
		return this.interval
	}

