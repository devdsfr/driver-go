/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { DriverServiceService } from './driver-service.service';

describe('Service: DriverService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [DriverServiceService]
    });
  });

  it('should ...', inject([DriverServiceService], (service: DriverServiceService) => {
    expect(service).toBeTruthy();
  }));
});
