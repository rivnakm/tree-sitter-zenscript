import XCTest
import SwiftTreeSitter
import TreeSitterZenscript

final class TreeSitterZenscriptTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_zenscript())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Zenscript grammar")
    }
}
