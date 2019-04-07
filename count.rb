require 'json'
require 'date'
require 'uri'
require 'pry'


def grab_all_scrapes(dir)
  scrape_file_regex = /scrape_\d\d\d\d\d\d\d\d_\d\d\d\d\d\d.json/
  scrapes = []
  Dir.foreach dir do |file|
    next unless file =~ scrape_file_regex

    str = File.read(File.join(dir, file))
    next if str.empty?

    h = JSON.parse(str, symbolize_names: true)
    scrapes << h
  end
  scrapes
end

def get_uniq_headlines(scrapes)
  uniq_headlines = {}

  scrapes.each do |scrape|
    headlines = scrape[:headlines]
    timeUnix = scrape[:scrapeTime]
    time = DateTime.strptime(timeUnix,'%s')

    headlines.each do |headline|
      url = headline[:url]
      title = headline[:title]
      blurb = headline[:blurb]
      uniq_headlines[url] = {} unless uniq_headlines[url]
      uniq_headlines[url][title] = [] unless uniq_headlines[url][title]
      uniq_headlines[url][title] << { time: time, blurb: blurb }
    end
  end

  uniq_headlines
end

dir = "/Users/mjohnsey/Downloads/wapo"
scrapes = grab_all_scrapes(dir)

uniq_headlines = get_uniq_headlines(scrapes)

uniq_headlines.each do |url, _occurrences|
  uri = URI(url)
  next if uri.host != "www.washingtonpost.com"
end

# uniq_headlines.select { |_, v| v.length > 1 }.each do |url, titles|
#   puts url
#   earliest_mention_per_title = titles.map {|title, instances| instances.map{|instance| instance[:time] }}.min
#   earliest_mention = titles.map {|_, instances| instances.map{|instance| instance[:time] }.min}.min
#   puts "First mentioned: #{earliest_mention}"
#   titles.each do |title, instances|
#     earliest = instances.map{|instance| instance[:time] }.min
#     latest = instances.map{|instance| instance[:time] }.max
#     puts "#{title} (#{earliest} - #{latest})"
#   end
#   puts "\n"
#   puts "="*5
# end
