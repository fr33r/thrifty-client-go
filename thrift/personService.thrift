namespace java com.thrifty.gen
namespace go thrifty

include "person.thrift"

/**
 * Service responsible for all interactions with Person objects.
 */
service PersonService {
  person.Person get(1: i32 id)
  i32 create(1: person.Person person)
  void remove(1: i32 id)
}
