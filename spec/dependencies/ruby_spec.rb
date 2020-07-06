describe 'Dependencies' do
  describe 'Ruby', dependency: true, ruby: true do
    versions, latest = versions_for :ruby

    versions.each do |version|
      describe "v#{version}" do
        around(:each) do |example|
          with_payload(deps: {'ruby' => version}) do |root|
            @root = root
            loon %w(up)

            assert_stderr_empty
            assert_status 0
            example.run
          end
        end

        it "installs correctly" do
          loon %w(exec ruby --version)
          assert_stdout version
        end

        it "sets the gem environment up" do
          loon %w(exec env)
          assert_stdout "GEM_HOME=#{@root}/.loon/data/gem"
          assert_stdout "GEM_PATH=#{@root}/.loon/data/gem"
          assert_stdout %r(PATH=.*#{@root}/.loon/data/gem/bin)
        end

        it "installs gems in the right location" do
          loon %w(exec gem install rake)
          assert_path "#{@root}/.loon/data/gem/bin/rake"
        end

        it "installs bundler" do
          assert_path "#{@root}/.loon/data/gem/bin/bundle"
          assert_path "#{@root}/.loon/data/gem/bin/bundler"
        end
      end
    end

    it "installs v#{latest} as the default" do
      test_dep 'ruby', match: latest
    end
  end
end
