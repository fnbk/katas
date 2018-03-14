require "test/unit"
require "./lib/reservation_service"

class ReservationServiceTest < Test::Unit::TestCase

  def setup
    @service = ReservationService.new
    @service.reset
  end

  def test_reserving_two_seats_in_empty_carriage
    reserved_seat_numbers = @service.reserve(2)
    assert_equal [1,2], reserved_seat_numbers
  end

  def test_reserving_three_seats_in_empty_carriage
    reserved_seat_numbers = @service.reserve(3)
    assert_equal [1,2,3], reserved_seat_numbers
  end

  def test_reserving_two_seats_in_non_empty_carriage
    @service.reserve(1)
    reserved_seat_numbers = @service.reserve(2)

    assert_equal [2,3], reserved_seat_numbers
  end

  def test_listing_when_no_reservations
    assert_equal [], @service.list_reserved(3)
  end

  def test_listing_reservations
    @service.reserve(1)
    @service.reserve(3)

    assert_equal [1,2,3], @service.list_reserved(3)
  end

end


describe ReservationService do
  #let(:service) { ReservationService.new }
  service =  ReservationService.new

  before do
    puts service.object_id
    service.reset
    #service.stub(:list_reserved).and_return("hello world")
  end

  context "empty carriage" do
    subject { service.reserve(number_of_seats) }
    let(:number_of_seats) { 2 }

    it "should reserve two seats" do
      subject.should eql [1,2]
    end

    context "three seats requested" do
      let(:number_of_seats) { 3 }

      it "should reserve three seats" do
        subject.should eql [1,2,3]
      end
    end

    it "should list no reservation" do
      service.list_reserved(3).should be_empty
    end
  end

  context "non-empty carriage" do
    subject { service }

    it "should reserve two seats" do
      subject.reserve(1)
      subject.reserve(2).should eql [2,3]
    end

    it "should list reservations" do
      subject.reserve(1)
      subject.reserve(3)
      subject.list_reserved(3).should eql [1,2,3]
    end
  end
end