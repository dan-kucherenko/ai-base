from rdflib import Graph, Namespace

def load_and_query_owl_file():
    # Load the OWL file
    ontology_file = "ontology.owl"
    g = Graph()
    g.parse(ontology_file, format="xml")

    # Print all triples in the graph
    print("All Triples in the Graph:")
    for triple in g:
        print(triple)

    # Define namespaces
    rdf = Namespace("http://www.w3.org/1999/02/22-rdf-syntax-ns#")
    xsd = Namespace("http://www.w3.org/2001/XMLSchema#")
    owl = Namespace("http://www.w3.org/2002/07/owl#")
    example = Namespace("http://example.org/")

    # Query for individuals and their ages
    query = """
        SELECT ?person ?age
        WHERE {
            ?person example:hasAge ?age.
        }
    """

    results = g.query(query, initNs={"example": example})

    print("\nIndividuals and their ages:")
    for row in results:
        print(f"{row['person']} has age {row['age']}")

    # Query for friendships
    query = """
        SELECT ?person1 ?person2
        WHERE {
            ?person1 example:hasFriend ?person2.
        }
    """

    results = g.query(query, initNs={"example": example})

    print("\nFriendships:")
    for row in results:
        print(f"{row['person1']} is friends with {row['person2']}")


if __name__ == "__main__":
    load_and_query_owl_file()
