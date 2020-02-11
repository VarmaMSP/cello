const TYPE_TEXT = 'T'
const TYPE_PODCAST = 'P'

const ICON_SEARCH = 'S'
const ICON_HISTORY = 'H'

export class SearchSuggestion {
  t: string
  i: string
  header: string
  subHeader: string

  static isTextSearch(s: SearchSuggestion): boolean {
    return s.t === TYPE_TEXT && s.i === ICON_SEARCH
  }

  static isTextHistory(s: SearchSuggestion): boolean {
    return s.t === TYPE_TEXT && s.i === ICON_HISTORY
  }

  static isPodcast(s: SearchSuggestion): boolean {
    return s.t === TYPE_PODCAST
  }

  constructor(j: any) {
    this.t = j['t'] || TYPE_TEXT
    this.i = j['i'] || ICON_SEARCH
    this.header = j['h1'] || ''
    this.subHeader = j['h2'] || ''
  }
}
