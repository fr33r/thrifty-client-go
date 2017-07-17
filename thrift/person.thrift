namespace java com.thrifty.gen
namespace go thrifty

include "address.thrift"

/**
 * Represents a human being.
 */
struct Person {
  1: string givenName,
  2: string surname,
  3: i8 age,
  4: address.Address address
}
