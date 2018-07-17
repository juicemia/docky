namespace :templates do
  task :generate do
    template_vars = {
      templates: {
        root: File.read('./templates/root.tpl.html'),
        route: File.read('./templates/routes.tpl.html'),
        index: File.read('./templates/index.tpl.html')
      }
    }

    src = File.join Dir.pwd, 'templates.go.erb'
    dst = File.join Dir.pwd, 'cmd/templates.go'
    execute_and_write_template src, template_vars, dst
  end

  # so nobody deletes templates.go.erb by accident
  task :clean do
    `rm templates.go`
  end
end

def execute_template(path, vars)
  require 'erb'
  require 'ostruct'

  ns = OpenStruct.new(vars)
  ERB.new(File.read(path)).result(ns.instance_eval { binding })
end

def execute_and_write_template(path, vars, target = File.join(File.dirname(path), File.basename(path, '.erb')))
  File.open(target, 'w+') do |f|
    puts "Evaluating ERB template #{path}..."
    output = execute_template(path, vars)

    puts "Writing result to #{target}..."
    f.write(output)
  end
end
