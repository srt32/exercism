class School

  def initialize
    @db = Hash.new()
  end

  def db
    @db
  end

  def add(name, grade)
    if @db[grade].nil?
      @db[grade] = []
    end
    @db[grade].push(name)
  end

  def grade(class_num)
    if @db[class_num].nil?
      @db[class_num] = []
    end
    @db[class_num]
  end

  def sort
    @db.sort_by{ |k,v| k}
  end

end
