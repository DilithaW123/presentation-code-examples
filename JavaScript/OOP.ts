abstract class Segment {
  hasTravelled: boolean = false;
  tattoo: number;
  constructor(tat: number) {
    this.tattoo = tat;
  }
  abstract logVendor(): void;
  abstract log(): void;
}

class AirSegment extends Segment {
  constructor(tat: number) {
    super(tat);
  }

  log(): void {
    console.log("Air Segment");
  }

  logVendor(): void {
    console.log("Air New Zealand");
  }
}

class CarSegment extends Segment {
  constructor(tat: number) {
    super(tat);
  }

  log(): void {
    console.log("Car Segment");
  }

  logVendor(): void {
    console.log("Hertz");
  }
}

class HotelSegment extends Segment {
  constructor(tat: number) {
    super(tat);
  }

  log(): void {
    console.log("Hotel Segment");
  }

  logVendor(): void {
    console.log("Hilton");
  }
}

function main() {
  let segments: Segment[] = [
    new AirSegment(1),
    new CarSegment(2),
    new HotelSegment(3),
  ];

  segments.forEach((segment) => {
    segment.log();
    segment.logVendor();
  });
}
