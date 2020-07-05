describe 'Dependencies' do
  describe 'Golang', dependency: true, golang: true do
    def test_golang_dep(version: nil, match:)
      dep = if version
              {'golang' => version}
            else
              'golang'
            end

      with_payload(deps: dep) do |project|
        loon %w(up), dir: project
        loon %w(exec go version), dir: project

        assert_stderr_empty
        assert_stdout match
        assert_status 0
      end
    end

    %w(1.13.12 1.14.4).each do |version|
      it "installs version #{version} correctly" do
        test_golang_dep version: version, match: version
      end
    end

    it 'installs version 1.14.4 as the default' do
      test_golang_dep match: '1.14.4'
    end
  end
end
