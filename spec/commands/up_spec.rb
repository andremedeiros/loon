describe 'Commands' do
  describe 'up', command: true do
    it "should bring up a project's infrastructure" do
      with_payload(deps: %w(memcached redis)) do |root|
        loon %w(up)
        assert_path "#{root}/.loon/pids/memcached.pid"
        assert_path "#{root}/.loon/pids/redis.pid"
      end
    end

    it 'should create an IP alias' do
      loon %(up)
      ip = project_ip

      ifconfig = if OS.linux?
        `ip addr list lo`
      else
        `ifconfig lo0`
      end

      expect(ifconfig).to include(ip)
    end
  end
end
